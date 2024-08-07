package service

import (
	"github.com/SoonDubu923/go-forum/dao/mysql"
	"github.com/SoonDubu923/go-forum/model"
	"github.com/SoonDubu923/go-forum/pkg/jwt"
	"github.com/SoonDubu923/go-forum/pkg/snowflake"
	"go.uber.org/zap"
)

// Register registers a new user.
func Register(p *model.ParamRegister) (err error) {
    // check if the username already exists
    if err = mysql.CheckIfUserExists(p.Username); err != nil {
        zap.L().Error("mysql.CheckIfUserExists failed", zap.Error(err))
        return
    }

    // generate a unique user ID
    userID := snowflake.GenID()

    // create a new user instance
    user := model.User{
        UserID:   userID,
        Username: p.Username,
        Password: p.Password,
    }

    // save the user to the database
    if err = mysql.SaveUser(&user); err != nil {
        zap.L().Error("mysql.SaveUser failed", zap.Error(err))
        return
    }

    return
}

// Login logs in a user.
func Login(p *model.ParamLogin) (tokenString string, err error) {
    // create a new user instance
    user := &model.User{
        Username: p.Username,
        Password: p.Password,
    }

    // log in the user
    if err = mysql.Login(user); err != nil {
        zap.L().Error("mysql.Login failed", zap.Error(err))
        return
    }

    // generate a JWT token
    return jwt.GenToken(user.UserID)
}