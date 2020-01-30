package logging

import (
    "github.com/sirupsen/logrus"
    "github.com/onrik/logrus/filename"
)

var log = logrus.New()

func New() *logrus.Logger {
    fHook := filename.NewHook()
    log.AddHook(fHook)
    return log
}
