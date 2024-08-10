package controller

import (
	"strconv"

	"github.com/SoonDubu923/go-forum/model"
	"github.com/SoonDubu923/go-forum/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

// PublishHandler handles the publish post request.
func PublishHandler(c *gin.Context) {
    // bind request parameters
    p := &model.Post{}
    if err := c.ShouldBindJSON(p); err != nil {
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
        zap.L().Error("invalid request parameters for PublishHandler", zap.Strings("errors", validationErrors))
        ErrorResponseWithMessage(c, CodeInvalidRequest, validationErrors)
        return
    }
    // get user ID from context (coming from AuthMiddleware)
    p.AuthorID = GetUser(c)

    // hand over to service layer
    if err := service.Publish(p); err != nil {
        zap.L().Error("service.Publish failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }

    SuccessResponse(c, CodeCreated, nil)
}

// PostDetailHandler handles the request for post details.
func PostDetailHandler(c *gin.Context) {
    // get post ID from URL
    postID ,err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        ErrorResponse(c, CodeInvalidParam)
        return
    }

    // hand over to service layer
    data, err := service.GetPostDetail(postID)
    if err != nil {
        zap.L().Error("service.GetPostDetail failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }

    // return response
    SuccessResponse(c, CodeSuccess, data)
}

// PostListHandler handles the request for post list.
func PostListHandler(c *gin.Context) {
    // get page info
    pageNum, pageSize, err := GetPageInfo(c)
    if err != nil {
        ErrorResponse(c, CodeInvalidParam)
        return
    }

    // hand over to service layer
    data, err := service.GetPostList(pageNum, pageSize)
    if err != nil {
        zap.L().Error("service.GetPostList failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }

    SuccessResponse(c, CodeSuccess, data)
}

// PostListHandlerUpdated handles the request for post list with ordering options.
// Order options: time, score.
func PostListHandlerUpdated(c *gin.Context) {
    // bind request parameters
    p := &model.ParamPostList{}

    if err := c.ShouldBindQuery(&p); err != nil {
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
        zap.L().Error("invalid request parameters for PostListHandlerUpdated", zap.Strings("errors", validationErrors))
        ErrorResponseWithMessage(c, CodeInvalidRequest, validationErrors)
        return
    }

    // hand over to service layer
    data, err := service.GetPostListUpdated(p)
    if err != nil {
        zap.L().Error("service.GetPostListUpdated failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }

    SuccessResponse(c, CodeSuccess, data)
}

func CommunityPostListHandler(c *gin.Context) {
    // bind request parameters
    p := &model.ParamCommunityPostList{}

    if err := c.ShouldBindQuery(&p); err != nil {
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
        zap.L().Error("invalid request parameters for CommunityPostListHandler", zap.Strings("errors", validationErrors))
        ErrorResponseWithMessage(c, CodeInvalidRequest, validationErrors)
        return
    }

    // hand over to service layer
    data, err := service.GetCommunityPostList(p)
    if err != nil {
        zap.L().Error("service.GetCommunityPostList failed", zap.Error(err))
        ErrorResponse(c, CodeServerError)
        return
    }

    SuccessResponse(c, CodeSuccess, data)
}