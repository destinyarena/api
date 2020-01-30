package discord

import (
    "fmt"
    "net/url"
    "net/http"
    "github.com/labstack/echo/v4"
    "crypto/rand"
    "encoding/hex"
)

/*
This is to prevent Cross Site requests by generating a state value that is passed in the url params
*/


func urlsafe(unsafeurl string) string {
    return url.QueryEscape(unsafeurl)
}


func randomHex(n int) (string, error) {
  bytes := make([]byte, n)
  if _, err := rand.Read(bytes); err != nil {
    return "", err
  }
  return hex.EncodeToString(bytes), nil
}

func GetAuthURL(c echo.Context) error {
    stateCode, _ := randomHex(10)
    fullauthurl := fmt.Sprintf("%s&state=%s", authurl, stateCode)
    return c.String(http.StatusOK, fullauthurl)
}

