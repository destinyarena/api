package discord

import (
    "fmt"
    "errors"
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

func getUser(p *RespOAuthPayload) (User, error) {
    client := new(http.Client)
    authtoken := fmt.Sprintf("%s %s", p.TokenType, p.AccessToken)
    log.Infoln(authtoken)

    req, err := http.NewRequest("GET", fmt.Sprintf("%s/v6/users/@me", cfg.BaseURL), nil)
    if err != nil {
        return User{}, err
    }

    req.Header.Set("Authorization", authtoken)
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        log.Error(err)
        return User{}, err
    }

    defer resp.Body.Close()

    userbody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return User{}, err
    }

    var user User
    json.Unmarshal(userbody, &user)

    return user, nil
}

func getToken(p *ReqPayload) (RespOAuthPayload, error) {
    client := new(http.Client)
    OAuth2URL := fmt.Sprintf("%s/oauth2/token?grant_type=%s&code=%s&redirect_uri=%s&scope=%s", cfg.BaseURL, "authorization_code", p.Code, urlsafe(cfg.RedirectURI), cfg.Scope)

    req, err := http.NewRequest("POST", OAuth2URL, nil)
    if err != nil {
        log.Error(err)
        return RespOAuthPayload{}, err
    }

    creds := b64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + secrets.Discord))
    req.Header.Set("Authorization", fmt.Sprintf("Basic %s", creds))
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err := client.Do(req)
    if err != nil {
        log.Error(err)
        return RespOAuthPayload{}, err
    }

    defer resp.Body.Close()

    authBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return RespOAuthPayload{}, err
    }

    if resp.StatusCode != 200 {
        return RespOAuthPayload{}, errors.New("Invalid error code")
    }

    var authPayload RespOAuthPayload
    json.Unmarshal(authBody, &authPayload)

    return authPayload, nil
}

func Callback(c echo.Context) (err error) {
    payload := new(ReqPayload)
    if err = c.Bind(payload); err != nil {
        return c.String(http.StatusBadRequest, "Invalid code payload")
    }

    v := validator.New()
    if errs := v.Struct(payload); errs != nil {
        return c.String(http.StatusBadRequest, "Invalid Payload")
    }

    authPayload, err := getToken(payload)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Well rip discord")
    }

    user, err := getUser(&authPayload)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Fuck me")
    }

    return c.JSON(http.StatusOK, user)
}
