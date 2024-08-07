package controller

import (
	"net/http"

	"github.com/SoonDubu923/go-forum/model"
	"github.com/SoonDubu923/go-forum/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
    // bind request parameters
    var p model.ParamRegister
    if err := c.ShouldBindJSON(&p); err != nil {
        var validationErrors []string
        // collect validation errors, if any
        if errs, ok := err.(validator.ValidationErrors); ok {
            for _, err := range errs {
                validationErrors = append(validationErrors, err.Error())
            }
        // if not a validation error, just append the error message
        } else {
            validationErrors = append(validationErrors, err.Error())
        }
        zap.L().Error("invalid request parameters for RegisterHandler", zap.Strings("errors", validationErrors))
        c.JSON(http.StatusBadRequest, gin.H{
            "status": "Error",
            "errors": validationErrors,
        })
        return
    }
    // hand over to service layer
    if err := service.Register(&p); err != nil {
        zap.L().Error("service.Register failed", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "Error",
            "message": "Registration failed",
        })
        return
    }
    // return success
    c.JSON(http.StatusOK, gin.H{
        "status": "Success",
        "message": "Registration successful",
    })
}