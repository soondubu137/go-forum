package controller

import "net/http"

type ResponseCode int

const (
    CodeSuccess ResponseCode = 600 + iota
    CodeCreated
    CodeServerError
    CodeInvalidRequest
    CodeInvalidParam
    CodeUserExists
    CodeUserNotExists
    CodeInvalidCredentials
    CodeTokenMissing
    CodeInvalidToken
    CodeServerBusy
    CodeNotFound
)

var codeMessageMap = map[ResponseCode]string{
    CodeSuccess:            "Success",
    CodeCreated:            "Created",
    CodeServerError:        "Unknown error",
    CodeInvalidRequest:     "Invalid request",
    CodeInvalidParam:       "Invalid parameter",
    CodeUserExists:         "User already exists",
    CodeUserNotExists:      "User does not exist",
    CodeInvalidCredentials: "Incorrect password",
    CodeTokenMissing:       "Token missing",
    CodeInvalidToken:       "Invalid token",
    CodeServerBusy:         "Server busy",
    CodeNotFound:           "Page not found",
}

var codeHttpStatusMap = map[ResponseCode]int{
    CodeSuccess:            http.StatusOK,
    CodeCreated:            http.StatusCreated,
    CodeServerError:        http.StatusInternalServerError,
    CodeInvalidRequest:     http.StatusBadRequest,
    CodeInvalidParam:       http.StatusBadRequest,
    CodeUserExists:         http.StatusConflict,
    CodeUserNotExists:      http.StatusNotFound,
    CodeInvalidCredentials: http.StatusUnauthorized,
    CodeTokenMissing:       http.StatusUnauthorized,
    CodeInvalidToken:       http.StatusUnauthorized,
    CodeServerBusy:         http.StatusServiceUnavailable,
    CodeNotFound:           http.StatusNotFound,
}

func (c ResponseCode) Message() string {
    return codeMessageMap[c]
}

func (c ResponseCode) HttpStatus() int {
    return codeHttpStatusMap[c]
}