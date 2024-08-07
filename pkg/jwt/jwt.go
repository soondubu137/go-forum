package jwt

import (
	"errors"
	"time"

	errmsg "github.com/SoonDubu923/go-forum/errors"

	"github.com/dgrijalva/jwt-go"
)

const ExpireDuration = time.Hour * 24

var secret = []byte("this is not a secret")

type CustomClaims struct {
    UserID   int64  `json:"user_id"`
    jwt.StandardClaims
}

// GenToken generates a new JWT token.
func GenToken(userID int64) (string, error) {
    c := CustomClaims{
        userID,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(ExpireDuration).Unix(),
            Issuer:    "go-forum",
        },
    }

    // Use HMAC for simplicity. You can use RSA or other methods for better security.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
    return token.SignedString(secret)
}

// ParseToken parses a JWT token.
func ParseToken(tokenString string) (*CustomClaims, error) {
    // parse and verify the token
    // note that token.Valid is populated here
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secret, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New(errmsg.ErrInvalidToken)
}