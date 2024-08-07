package routes

import (
	"net/http"

	"github.com/SoonDubu923/go-forum/controller"
	"github.com/SoonDubu923/go-forum/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
    r := gin.New()
    r.Use(logger.GinLogger(), logger.GinRecovery(true))

    // register routes here
    r.POST("/register", controller.RegisterHandler)

    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"status": "Error", "message": "Page not found"})
    })
    return r
}