package middleware

import (
    "net/http"
    "github.com/labstack/echo/middlware"
    "github.com/arturoguerra/destinyarena-api/internal/config"
)

var secret = config.LoadSecrets()
var BotAuth = middleware.JWTWithConfig(middleware.JWTConfig{
    SigningKey: []byte(secret.JWTSecret),
}
