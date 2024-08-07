package controller

import (
	"github.com/SoonDubu923/go-forum/model"
	"github.com/SoonDubu923/go-forum/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	errmsg "github.com/SoonDubu923/go-forum/errors"
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
        ErrorResponseWithMessage(c, CodeInvalidRequest, validationErrors)
        return
    }
    // hand over to service layer
    if err := service.Register(&p); err != nil {
        zap.L().Error("service.Register failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }
    // return success
    SuccessResponse(c, CodeSuccess, nil)
}

func LoginHandler(c *gin.Context) {
    // bind request parameters
    var p model.ParamLogin
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
        zap.L().Error("invalid request parameters for LoginHandler", zap.Strings("errors", validationErrors))
        ErrorResponseWithMessage(c, CodeInvalidRequest, validationErrors)
        return
    }
    // hand over to service layer
    tokenString, err := service.Login(&p)
    if err != nil {
        zap.L().Error("service.Login failed", zap.Error(err))
        switch err.Error() {
        case errmsg.ErrIncorrectCredentials:
            ErrorResponseWithMessage(c, CodeInvalidCredentials, errmsg.ErrIncorrectCredentials)
        case errmsg.ErrInvalidToken:
            ErrorResponseWithMessage(c, CodeInvalidToken, errmsg.ErrInvalidToken)
        default:
            ErrorResponse(c, CodeServerError)
        }
        return
    }
    // return success
    SuccessResponse(c, CodeSuccess, tokenString)
}