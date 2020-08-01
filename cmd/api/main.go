package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/nats-io/nats.go"


    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
    "github.com/arturoguerra/destinyarena-api/internal/logging"

    "github.com/arturoguerra/destinyarena-api/internal/router/oauth"
    "github.com/arturoguerra/destinyarena-api/internal/router/registration"
    "github.com/arturoguerra/destinyarena-api/internal/router/users"
)

const (
    STANCLIENT = "arena-api"
    DISCORD_REGISTRATION = "registration"
)

func main() {
    log := logging.New()

    // NATS Client setup
    ncfg := config.LoadNATSConfig()
    log.Infof("Starting NATS Client: %s", ncfg.URL)
    nc, err := nats.Connect(ncfg.URL)
    if err != nil {
        log.Fatal(err)
    }

    ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

    sendRegistrationChan := make(chan *structs.NATSRegistration)
    ec.BindSendChan(DISCORD_REGISTRATION, sendRegistrationChan)

    nchan := &structs.NATS{
        SendRegistration: sendRegistrationChan,
    }

    // Echo Server Setup
    log.Infoln("Starting Destiny Arena API")
    e := echo.New()

    // Recover Application from a panic anywhere in the chain
    e.Use(middleware.Recover())

    g := e.Group("/api", middleware.Logger())

    oauth.New(g)
    users.New(g)
    registration.New(g, nchan)

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "destinyarena api base")
    })

    host, port := config.LoadHTTPConfig()
    log.Infof("Running with HOST: %s PORT: %s", host, port)

    e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", host, port)))
}
