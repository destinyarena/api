package oauth

import (
    "github.com/arturoguerra/destinyarena-api/internal/router/oauth/discord"
    "github.com/arturoguerra/destinyarena-api/internal/router/oauth/bungie"
    "github.com/arturoguerra/destinyarena-api/internal/router/oauth/faceit"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    "github.com/labstack/echo/v4"
)

func New(e *echo.Group) {
    logger := logging.New()
    logger.Info("Registring Group /api/oauth")

    g := e.Group("/oauth")
    discord.New(g)
    faceit.New(g)
    bungie.New(g)
}
