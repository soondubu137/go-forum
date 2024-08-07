package routes

import (
	"github.com/SoonDubu923/go-forum/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
    r := gin.New()
    r.Use(logger.GinLogger(), logger.GinRecovery(true))
    // register routes here
    return r
}