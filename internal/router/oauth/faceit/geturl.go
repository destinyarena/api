package faceit

import (
    "fmt"
    "net/http"
    "github.com/labstack/echo/v4"
)

func GetOAuthURL(c echo.Context) error {
    authURL := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_popup=true", cfg.OAuthURL, cfg.ClientID)
    return c.String(http.StatusOK, authURL)
}
