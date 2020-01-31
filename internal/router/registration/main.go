package registration

import (
    "github.com/sirupsen/logrus"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/pkg/database"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
)

var log *logrus.Logger
var dbclient *database.DBClient

func init() {
    log = logging.New()
}

func New(e *echo.Echo, client *database.DBClient) {
    log.Infoln("Registering POST /api/registration")
    e.POST("/api/registration", endpoint)
}
