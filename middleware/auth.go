package middleware

import (
	"strings"

	"github.com/SoonDubu923/go-forum/controller"
	"github.com/SoonDubu923/go-forum/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware that checks if the request is authenticated.
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // get the authorization header
        authHeader := c.GetHeader("Authorization")

        // if the authorization header is empty, return an error
        if authHeader == "" {
            controller.ErrorResponseWithMessage(c, controller.CodeTokenMissing, "Authorization header is empty")
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        // check if the authorization header format is correct
        if len(parts) != 2 || parts[0] != "Bearer" {
            controller.ErrorResponseWithMessage(c, controller.CodeTokenMissing, "Authorization header format must be Bearer {token}")
            c.Abort()
            return
        }

        // parse and validate the token
        token, err := jwt.ParseToken(parts[1])
        if err != nil {
            controller.ErrorResponse(c, controller.CodeInvalidToken)
            c.Abort()
            return
        }

        // set the user ID in the context and continue
        c.Set(controller.USER_ID, token.UserID)
        c.Next()
    }
}