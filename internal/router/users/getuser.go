package users

import (
    "net/http"
    "github.com/labstack/echo/v4"
    faceitApi "github.com/arturoguerra/destinyarena-api/pkg/faceit"
)

type (
    User struct {
        Discord   string `json:"discord"   validate:"required"`
        Bungie    string `json:"bungie"    validate:"required"`
        Faceit    string `json:"faceit"    validate:"required"`
        FaceitLvl int    `json:"faceitlvl" validate:"required"`
    }
)

func getLvl(id string) (int, error) {
    client := faceitApi.New(secrets.FaceitAPIKey)
    user, err := client.GetUser(id)
    if err != nil {
        log.Error(err)
        return 0, err
    }

    return user.SkillLevel, nil
}

func getUser(c echo.Context) error {
    id := c.Param("id")
    err, dbuser := dbclient.GetUser(id)
    if err != nil {
        log.Error(err)
        return c.String(404, "User is not alive")
    }

    lvl, err := getLvl(dbuser.Faceit)
    if err != nil {
        log.Error(err)
    }

    user := &User{
        Discord: dbuser.Discord,
        Bungie: dbuser.Bungie,
        Faceit: dbuser.Faceit,
        Banned: dbuser.Banned,
        FaceitLvl: lvl,
    }

    return c.JSON(http.StatusOK, user)
}
