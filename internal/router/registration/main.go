package registration

import (
    "github.com/surupsen/logrus"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/logging"
)

var log = logrus.Logger

func int() {
    log = logging.New()

func New(e *echo.Echo) {
    log.Infoln("Registering POST /api/registration")
    e.POST("/api/registration", endpoint)
}
