package middleware

import (
    "strings"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
)

var secret = config.LoadSecrets()
var log = logging.New()

func BotAuth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        log.Debugln(token)
        if token == "" {
            return c.String(401, "Invalid Auth token")
        }

        if strings.HasPrefix(token, "Bearer ") {
            trueToken := strings.TrimPrefix(token, "Bearer ")
            if trueToken == secret.APIKey {
                return next(c)
            }
        }

        return c.String(401, "Invalid Auth token")
    }
}
