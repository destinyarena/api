package main

import (
    "os"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/arturoguerra/destinyarena-api/internal/router"
    "github.com/arturoguerra/destinyarena-api/pkg/database"
)

func main() {
    e := echo.New()

    // Logs HTTP Requests
    e.Use(middleware.Logger())
    // Recover Application from a panic anywhere in the chain
    e.Use(middleware.Recover())

    // Initial Database connection
    username := ""
    password := ""
    host := ""
    err, db := database.New(username, password, host)
    if err != nil {
        panic(err)
    }

    // New router in its own thread
    go func() {
        router.New(e, db)
    }()

    // HTTP Server PORT Defaults to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    e.Logger.Fatal(e.Start(fmt.Sprinf(":%s", port)))

}
