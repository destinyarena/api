package main

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"


    "github.com/arturoguerra/destinyarena-api/internal/config"
    "github.com/arturoguerra/destinyarena-api/internal/logging"

    "github.com/arturoguerra/destinyarena-api/internal/router/oauth"
    "github.com/arturoguerra/destinyarena-api/internal/router/registration"
)

func main() {
    log := logging.New()
    log.Infoln("Starting Destiny Arena API")
    e := echo.New()

    // Logs HTTP Requests
    e.Use(middleware.Logger())
    // Recover Application from a panic anywhere in the chain
    e.Use(middleware.Recover())

    g := e.Group("/api", middleware.Logger())

    go oauth.New(g)
    go registration.New(g)

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "destinyarena api base")
    })

    host, port := config.LoadHTTPConfig()
    log.Infof("Running with HOST: %s PORT: %s", host, port)

    e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", host, port)))

}
