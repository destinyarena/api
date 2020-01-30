package faceit

import (
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
    "github.com/labstack/echo/v4"
)

var cfg *structs.Faceit

func init() {
    cfg = config.LoadFaceitConfig()
}

func New(g *echo.Group) {
    g.GET("/faceit/authurl", GetOAuthURL)
    g.GET("/faceit/callback", Callback)
}
