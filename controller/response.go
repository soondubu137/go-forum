package controller

import "github.com/gin-gonic/gin"

type Response struct {
    Code    ResponseCode `json:"code"`
    Status  string       `json:"status"`
    Message any          `json:"message,omitempty"`
    Data    any          `json:"data,omitempty"`
}

func ErrorResponse(c *gin.Context, code ResponseCode) {
    c.JSON(code.HttpStatus(), Response{
        Code:    code,
        Status:  "error",
        Message: code.Message(),
        Data:    nil,
    })
}

func ErrorResponseWithMessage(c *gin.Context, code ResponseCode, message any) {
    c.JSON(code.HttpStatus(), Response{
        Code:    code,
        Status:  "error",
        Message: message,
        Data:    nil,
    })
}

func SuccessResponse(c *gin.Context, code ResponseCode, data any) {
    c.JSON(code.HttpStatus(), Response{
        Code:    code,
        Status:  "success",
        Message: code.Message(),
        Data:    data,
    })
}