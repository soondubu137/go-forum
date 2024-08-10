package controller

import (
	"github.com/SoonDubu923/go-forum/model"
	"github.com/SoonDubu923/go-forum/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func VoteHandler(c *gin.Context) {
    // bind request parameters
    var p model.ParamVoteData
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
        zap.L().Error("invalid request parameters for VoteHandler", zap.Strings("errors", validationErrors))
        ErrorResponseWithMessage(c, CodeInvalidRequest, validationErrors)
        return
    }

    // hand over to service layer
    userID := GetUser(c)
    if err := service.VoteForPost(userID, &p); err != nil {
        zap.L().Error("service.VoteForPost() failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }

    SuccessResponse(c, CodeSuccess, nil)
}