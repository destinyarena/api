package users

import (
    "fmt"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/labstack/echo/v4"
)

var secrets = config.LoadSecrets()

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        if token == fmt.Sprintf("Basic %s", secrets.APIKey) {
            return next(c)
        } else {
            return c.String(403, "Invalid token")
        }
    }
}
