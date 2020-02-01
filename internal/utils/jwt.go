package utils

import (
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/dgrijalva/jwt-go"

    "errors"
    "fmt"
)

var secrets = config.LoadSecrets()

func SignJWT(claims jwt.Claims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    if secrets.JWTSecret == "" {
        err := errors.New("Invalid Secret")
        return "", err
    }

    tokenString, err := token.SignedString([]byte(secrets.JWTSecret))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func DecryptJWT(tokenString string, claims jwt.Claims) (error) {
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return []byte(secrets.JWTSecret), nil
    })

    if err != nil {
        return err
    }

    if !token.Valid {
        err = errors.New("Invalid token")
        return err
    }

    return nil
}
