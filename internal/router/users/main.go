package users

import (
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    "github.com/labstack/echo/v4"
)

var log = logging.New()
var grpcfg = config.LoadGRPConfig()

func New(g *echo.Group) {
    g.GET("/users/:id", GetId, Auth)
    g.GET("/users/list", List, Auth)
}
