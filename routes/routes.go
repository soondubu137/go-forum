package routes

import (
	"net/http"

	"github.com/SoonDubu923/go-forum/controller"
	"github.com/SoonDubu923/go-forum/logger"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
    // if in release mode, use the release mode of Gin
    if mode == gin.ReleaseMode {
        gin.SetMode(gin.ReleaseMode)
    }

    r := gin.New()
    r.Use(logger.GinLogger(), logger.GinRecovery(true))

    // register routes here
    r.POST("/register", controller.RegisterHandler)

    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"status": "Error", "message": "Page not found"})
    })
    return r
}