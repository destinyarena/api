package users

import (
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    "github.com/labstack/echo/v4"
)

var log = logging.New()

func New(g *echo.Group) {
    e.GET("/users/:id", getId, Auth)
    e.GET("/users/list", GetAll, Auth)
}
