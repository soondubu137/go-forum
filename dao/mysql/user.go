package mysql

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	errmsg "github.com/SoonDubu923/go-forum/errors"
	"github.com/SoonDubu923/go-forum/model"
)

// CheckIfUserExists checks if the provided username exists in the database.
func CheckIfUserExists(username string) (err error) {
    var count int
    if err = db.Get(&count, "SELECT count(id) FROM user WHERE username = ?", username); err != nil {
        return
    }
    if count > 0 {
        return errors.New(errmsg.ErrUserExists)
    }
    return
}

// SaveUser saves the provided user to the database.
func SaveUser(user *model.User) error {
    // encrypt the password
    if err := encryptPassword(user); err != nil {
        return err
    }
    _, err := db.NamedExec("INSERT INTO user (user_id, username, password, salt) VALUES (:user_id, :username, :password, :salt)", user)
    return err
}

func encryptPassword(user *model.User) error {
    // generate a random salt
    salt := make([]byte, 8)
    if _, err := rand.Read(salt); err != nil {
        return err
    }

    // concatenate the password and salt
    saltedPassword := append([]byte(user.Password), salt...)

    // hash the salted password
    hashedPassword := sha256.Sum256(saltedPassword)
    
    // save the hashed password and salt
    user.Password = hex.EncodeToString(hashedPassword[:])
    user.Salt = hex.EncodeToString(salt)

    return nil
}

// Login logs in a user.
func Login(p *model.User) error {
    var user model.User
    // retrieve the user from the database
    if err := db.Get(&user, "SELECT user_id, username, password, salt FROM user WHERE username = ?", p.Username); err != nil {
        if err.Error() == "sql: no rows in result set" {
            return errors.New(errmsg.ErrIncorrectCredentials)
        }
        return err
    }

    // concatenate the provided password and the salt from the database
    salt, err := hex.DecodeString(user.Salt)
    if err != nil {
        return err
    }
    saltedPassword := append([]byte(p.Password), salt...)

    // hash the salted password
    hashedPassword := sha256.Sum256(saltedPassword)
    if hex.EncodeToString(hashedPassword[:]) != user.Password {
        return errors.New(errmsg.ErrIncorrectCredentials)
    }
    return nil
}