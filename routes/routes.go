package routes

import (
	"scaffold/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
    r := gin.New()
    r.Use(logger.GinLogger(), logger.GinRecovery(true))
    // register routes
    return r
}