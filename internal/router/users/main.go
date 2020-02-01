package users

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/pkg/database"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/router/middleware"
)

var log = logging.New()
var dbclient *database.DBClient
var secrets = config.LoadSecrets()

func New(e *echo.Echo, client *database.DBClient) {
    dbclient = client

    g := e.Group("/api/users", middleware.BotAuth)

    g.GET("/get/:id", getUser)
}
