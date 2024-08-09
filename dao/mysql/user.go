package mysql

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"

	errmsg "github.com/SoonDubu923/go-forum/errors"
	"github.com/SoonDubu923/go-forum/model"
	"go.uber.org/zap"
)

// CheckIfUserExists checks if the provided username exists in the database.
func CheckIfUserExists(username string) (err error) {
    var count int
    if err = db.Get(&count, "SELECT count(id) FROM user WHERE username = ?", username); err != nil {
        zap.L().Error("CheckIfUserExists failed", zap.String("username", username), zap.Error(err))
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
        zap.L().Error("encryptPassword failed", zap.Error(err))
        return err
    }
    if _, err := db.NamedExec("INSERT INTO user (user_id, username, password, salt) VALUES (:user_id, :username, :password, :salt)", user); err != nil {
        zap.L().Error("SaveUser failed", zap.Error(err))
        return err
    }
    return nil
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
        if err == sql.ErrNoRows {
            return errors.New(errmsg.ErrIncorrectCredentials)
        }
        zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err))
        return err
    }

    // concatenate the provided password and the salt from the database
    salt, err := hex.DecodeString(user.Salt)
    if err != nil {
        return err
    }
    saltedPassword := append([]byte(p.Password), salt...)

    // hash the salted password and compare it with the hashed password from the database
    hashedPassword := sha256.Sum256(saltedPassword)
    if hex.EncodeToString(hashedPassword[:]) != user.Password {
        return errors.New(errmsg.ErrIncorrectCredentials)
    }

    // p is a pointer and the UserID field is needed in the service layer for JWT generation
    p.UserID = user.UserID
    return nil
}

// GetUsernameByID gets a username by user ID.
func GetUsernameByID(userID int64) (string, error) {
    var username string
    if err := db.Get(&username, "SELECT username FROM user WHERE user_id = ?", userID); err != nil {
        zap.L().Error("GetUsernameByID failed", zap.Int64("userID", userID), zap.Error(err))
        return "", err
    }
    return username, nil
}