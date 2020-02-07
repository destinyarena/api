package registration

import (
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/logging"
    pb "github.com/arturoguerra/destinyarena-api/pkg/profiles"
)

var (
    log = logging.New()
    uClient pb.ProfilesClient
)

func New(e *echo.Group) {
    log.Infoln("Registering POST /api/registration")
    e.POST("/registration", endpoint)

    grpcfg := config.LoadGRPConfig()
    uClient, err := pb.New(grpcfg.ProfilesHost, grpcfg.ProfilesPort)
    if err != nil {
        log.Error(err)
    }
    var _ = uClient
}
