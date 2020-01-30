package bungie

import (
    "fmt"
    "net/http"
    "crypto/rand"
    "encoding/hex"
    "github.com/labstack/echo/v4"
)

func randomHex(n int) (string, error) {
  bytes := make([]byte, n)
  if _, err := rand.Read(bytes); err != nil {
    return "", err
  }
  return hex.EncodeToString(bytes), nil
}


func GetOAuthURL(c echo.Context) error {
    stateCode, _ := randomHex(10)
    returnurl := fmt.Sprintf("%s?client_id=%s&response_type=code&state=%s", cfg.OAuth2URL, cfg.ClientID, stateCode)
    return c.String(http.StatusOK, returnurl)
}
