package controller

import (
	"strconv"

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

func GetPageInfo(c *gin.Context) (int64, int64, error) {
    pageNumStr := c.Query("pageNum")
    pageSizeStr := c.Query("pageSize")
    
    var (
        pageNum int64
        pageSize int64
        err error
    )

    pageNum, err = strconv.ParseInt(pageNumStr, 10, 32)
    if err != nil {
        return 0, 0, err
    }
    pageSize, err = strconv.ParseInt(pageSizeStr, 10, 32)
    if err != nil {
        return 0, 0, err
    }

    return pageNum, pageSize, nil
}