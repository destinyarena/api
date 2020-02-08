package registration

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
)

var log = logging.New()
var nchan *structs.NATS

func New(e *echo.Group, nc *structs.NATS ) {
    nchan = nc
    log.Infoln("Registering POST /api/registration")
    e.POST("/registration", endpoint)
}
