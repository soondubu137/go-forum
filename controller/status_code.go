package controller

import "net/http"

type ResponseCode int

const (
    CodeSuccess ResponseCode = 600 + iota
    CodeCreated
    CodeServerError
    CodeInvalidRequest
    CodeUserExists
    CodeUserNotExists
    CodeInvalidCredentials
    CodeTokenMissing
    CodeInvalidToken
    CodeServerBusy
)

var codeMessageMap = map[ResponseCode]string{
    CodeSuccess:            "Success",
    CodeCreated:            "Created",
    CodeServerError:        "Unknown error",
    CodeInvalidRequest:     "Invalid request",
    CodeUserExists:         "User already exists",
    CodeUserNotExists:      "User does not exist",
    CodeInvalidCredentials: "Incorrect password",
    CodeTokenMissing:       "Token missing",
    CodeInvalidToken:       "Invalid token",
    CodeServerBusy:         "Server busy",
}

var codeHttpStatusMap = map[ResponseCode]int{
    CodeSuccess:            http.StatusOK,
    CodeCreated:            http.StatusCreated,
    CodeServerError:        http.StatusInternalServerError,
    CodeInvalidRequest:     http.StatusBadRequest,
    CodeUserExists:         http.StatusConflict,
    CodeUserNotExists:      http.StatusNotFound,
    CodeInvalidCredentials: http.StatusUnauthorized,
    CodeTokenMissing:       http.StatusUnauthorized,
    CodeInvalidToken:       http.StatusUnauthorized,
    CodeServerBusy:         http.StatusServiceUnavailable,
}

func (c ResponseCode) Message() string {
    return codeMessageMap[c]
}

func (c ResponseCode) HttpStatus() int {
    return codeHttpStatusMap[c]
}