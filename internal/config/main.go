package config

import (
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    "github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
    log = logging.New()
}
