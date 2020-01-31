package discord

import (
    "github.com/labstack/echo/v4"
    "github.com/sirupsen/logrus"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
)
var (
    log *logrus.Logger
    secrets *structs.Secrets
    cfg *structs.Discord
)

func init() {
    log = logging.New()
    cfg = config.LoadDiscordConfig()
    secrets = config.LoadSecrets()

}

func New(g *echo.Group) {
    log.Infoln("Registering GET /api/oauth/discord/authurl")
    g.GET("/discord/authurl", GetAuthURL)

    log.Infoln("Registering GET /api/oauth/discord/callback")
    g.GET("/discord/callback", Callback)
}
