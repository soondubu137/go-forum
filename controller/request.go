package controller

import (
	"github.com/gin-gonic/gin"
)

const (
    USER_ID = "userID"
)

// GetUser returns the user ID from the context.
func GetUser(c *gin.Context) int64 {
    userID, _ := c.Get(USER_ID)
    return userID.(int64)
}