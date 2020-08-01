package users

import (
    "context"
    "github.com/arturoguerra/destinyarena-api/pkg/profiles"
    "google.golang.org/grpc"
    "github.com/labstack/echo/v4"
)


type Profile struct {
    Discord string `json:"discord"`
    Bungie  string `json:"bungie"`
    Faceit  string `json:"faceit"`
    Banned  bool   `json:"banned"`
}

func GetId(c echo.Context) error {
    id := c.Param("id")
    if id == "" {
        return c.String(404, "Invalid ID")
    }

    conn, err := grpc.Dial(grpcfg.Profiles, grpc.WithInsecure())
    if err != nil {
        log.Error(err)
        return c.String(502, err.Error())
    }

    defer conn.Close()

    p := profiles.NewProfilesClient(conn)
    log.Info("Getting user profile")
    r, err := p.GetProfile(context.Background(), &profiles.IdRequest{
        Id: id,
    })

    if err != nil {
        log.Error(err)
        return c.String(404, err.Error())
    }

    return c.JSON(200, &Profile{
        Discord: r.GetDiscord(),
        Bungie: r.GetBungie(),
        Faceit: r.GetFaceit(),
        Banned: r.GetBanned(),
    })
}
