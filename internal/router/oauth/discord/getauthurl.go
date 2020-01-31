package discord

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/arturoguerra/destinyarena-api/internal/utils"
)

/*
This is to prevent Cross Site requests by generating a state value that is passed in the url params
*/

func GetAuthURL(c echo.Context) error {
    stateCode, _ := utils.RandomHex(10)
    authurl := fmt.Sprintf("%s/oauth2/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s", cfg.BaseURL, cfg.ClientID, cfg.Scope, utils.SafeUrl(cfg.RedirectURI))
    fullauthurl := fmt.Sprintf("%s&state=%s", authurl, stateCode)
    return c.String(http.StatusOK, fullauthurl)
}

