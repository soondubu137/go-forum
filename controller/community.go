package controller

import (
	"github.com/SoonDubu923/go-forum/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
    data, err := service.GetCommunities()
    if err != nil {
        zap.L().Error("service.GetCommunities failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }
    SuccessResponse(c, CodeSuccess, data)
}