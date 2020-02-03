package invites

/*
The package gets invites from faceit"
*/

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    "github.com/arturoguerra/destinyarena-api/internal/router/middleware"
)

var secrets = config.LoadSecrets()
var log = logging.New()

func New(e *echo.Group) {
    g := e.Group("/invites", middleware.BotAuth)
    g.GET("/:id", getInvite)
}
