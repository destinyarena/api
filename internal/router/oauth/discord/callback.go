package discord

import (
    "fmt"
//    "bytes"
    b64 "encoding/base64"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/labstack/echo/v4"
    "gopkg.in/go-playground/validator.v9"
)


type (
    ReqPayload struct {
        Code string `json:"code" validate:"required"`
    }

    RespOAuthPayload struct {
        AccessToken  string `json:"access_token"`
        TokenType    string `json:"token_type"`
        ExpiresIn    string `json:"expires_in"`
        RefreshToken string `json:"refresh_token"`
        Scope        string `json:"scope"`
    }

    User struct {
        ID            string `json:"id"`
        Username      string `json:"username"`
        Discriminator string `json:"discriminator"`
    }
)

func Callback(c echo.Context) (err error) {
    payload := new(ReqPayload)
    if err = c.Bind(payload); err != nil {
        return c.String(http.StatusBadRequest, "Invalid code payload")
    }

    v := validator.New()
    if errs := v.Struct(payload); errs != nil {
        return c.String(http.StatusBadRequest, "Invalid Payload")
    }

    client := new(http.Client)


    OAuth2URL := fmt.Sprintf("%s/oauth2/token?grant_type=%s&code=%s&redirect_uri=%s&scope=%s", cfg.BaseURL, "authorization_code", payload.Code, urlsafe(cfg.RedirectURI), cfg.Scope)

    req, err := http.NewRequest("POST", OAuth2URL, nil)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Well shit we broken something lol")
    }

    creds := b64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + secrets.Discord))
    req.Header.Set("Authorization", fmt.Sprintf("Basic %s", creds))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err := client.Do(req)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Well shit something went wrong")
    }

    defer resp.Body.Close()

    authBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Looks like something went wrong with discord")
    }

    if resp.StatusCode != 200 {
        return c.String(http.StatusBadRequest, "Error with discord's payload")
    }

    var authPayload RespOAuthPayload
    json.Unmarshal(authBody, &authPayload)

    authtoken := fmt.Sprintf("%s %s", authPayload.TokenType, authPayload.AccessToken)
    log.Infoln(authtoken)

    userreq, err := http.NewRequest("GET", fmt.Sprintf("%s/v6/users/@me", cfg.BaseURL), nil)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Well shit we broken something lol")
    }

    userreq.Header.Set("Authorization", authtoken)
    userreq.Header.Set("Content-Type", "application/json")
    userresp, err := client.Do(userreq)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Something went wrong please try again later")
    }

    defer userresp.Body.Close()

    userbody, err := ioutil.ReadAll(userresp.Body)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "FUCK YOU")
    }

    var user User
    json.Unmarshal(userbody, &user)

    return c.JSON(http.StatusOK, user)
}
