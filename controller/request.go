package controller

import (
	"errors"

	errmsg "github.com/SoonDubu923/go-forum/errors"

	"github.com/gin-gonic/gin"
)

const (
    USER_ID = "userID"
)

// GetUser returns the user ID from the context.
func GetUser(c *gin.Context) (int64, error) {
    userID, ok := c.Get(USER_ID)
    if !ok {
        return 0, errors.New(errmsg.ErrUnauthenticated)
    }
    return userID.(int64), nil
}