package faceit

import (
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/labstack/echo/v4"
    "github.com/sirupsen/logrus"
)

var (
    log *logrus.Logger
    secrets *structs.Secrets
    cfg *structs.Faceit
)

func init() {
    log = logging.New()
    cfg = config.LoadFaceitConfig()
    secrets = config.LoadSecrets()
}

func New(g *echo.Group) {
    g.GET("/faceit/authurl", GetOAuthURL)
    g.GET("/faceit/callback", Callback)
}
