package mysql

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/SoonDubu923/go-forum/model"
)

// CheckIfUserExists checks if the provided username exists in the database.
func CheckIfUserExists(username string) (err error) {
    var count int
    if err = db.Get(&count, "SELECT count(id) FROM user WHERE username = ?", username); err != nil {
        return
    }
    if count > 0 {
        return errors.New("username already exists")
    }
    return
}

// SaveUser saves the provided user to the database.
func SaveUser(user *model.User) error {
    // encrypt the password
    encryptedPassword, salt, err := encryptPassword(user.Password)
    if err != nil {
        return err
    }
    _ , err = db.Exec("INSERT INTO user (user_id, username, password, salt) VALUES (?, ?, ?, ?)", user.UserID, user.Username, encryptedPassword, salt)
    return err
}

func encryptPassword(password string) (string, string, error) {
    // generate a random salt
    salt := make([]byte, 8)
    if _, err := rand.Read(salt); err != nil {
        return "", "", err
    }

    // concatenate the password and salt
    saltedPassword := append([]byte(password), salt...)

    // hash the salted password
    hashedPassword := sha256.Sum256(saltedPassword)
    return hex.EncodeToString(hashedPassword[:]), hex.EncodeToString(salt), nil
}
