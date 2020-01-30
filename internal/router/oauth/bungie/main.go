package bungie

import (
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
    "github.com/labstack/echo/v4"
    "github.com/sirupsen/logrus"
)

var (
    log *logrus.Logger
    secrets *structs.Secrets
    cfg *structs.Bungie
)

func init() {
    log = logging.New()
    cfg = config.LoadBungieConfig()
    secrets = config.LoadSecrets()
}

func New(g *echo.Group) {
    g.GET("/bungie/authurl", GetOAuthURL)
}
