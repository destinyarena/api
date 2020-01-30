package discord

import (
    "fmt"
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
    authurl string
)

func init() {
    log = logging.New()
    cfg = config.LoadDiscordConfig()
    secrets = config.LoadSecrets()
    authurl = fmt.Sprintf("%s/oauth2/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s", cfg.BaseURL, cfg.ClientID, cfg.Scope, urlsafe(cfg.RedirectURI))

}

func New(g *echo.Group) {
    log.Infoln("Registering GET /api/oauth/discord/authurl")
    g.GET("/discord/authurl", GetAuthURL)

    log.Infoln("Registering GET /api/oauth/discord/callback")
    g.GET("/discord/callback", Callback)
}
