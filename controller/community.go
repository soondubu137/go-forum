package controller

import (
	"strconv"

	errmsg "github.com/SoonDubu923/go-forum/errors"
	"github.com/SoonDubu923/go-forum/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CommunityHandler handles the request to get a list of communities.
func CommunityHandler(c *gin.Context) {
    data, err := service.GetCommunities()
    if err != nil {
        zap.L().Error("service.GetCommunities failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }
    SuccessResponse(c, CodeSuccess, data)
}

// CommunityDetailHandler handles the request to get the details of a community.
func CommunityDetailHandler(c *gin.Context) {
    // get the community ID from the URL
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        ErrorResponse(c, CodeInvalidParam)
        return
    }

    data, err := service.GetCommunityDetail(id)
    if err != nil {
        if err.Error() == errmsg.ErrNotFound {
            ErrorResponse(c, CodeNotFound)
            return
        }
        zap.L().Error("service.GetCommunityDetail failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }

    SuccessResponse(c, CodeSuccess, data)
}