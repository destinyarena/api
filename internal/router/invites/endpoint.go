package invites

import (
    "net/http"
    "github.com/arturoguerra/destinyarena-api/pkg/faceit"
    "github.com/labstack/echo/v4"
)

func getInvite(c echo.Context) error {
    hubid := c.Param("id")

    f := faceit.New(secrets.FaceitUserKey)

    err, link := f.GetInvite(hubid)
    if err != nil {
        // TODO: Stuff
        return c.String(403, "")
    }

    return c.String(http.StatusOK, link)
}
