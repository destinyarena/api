package registration

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
)

var log = logging.New()

func New(e *echo.Group) {
    log.Infoln("Registering POST /api/registration")
    e.POST("/registration", endpoint)
}
