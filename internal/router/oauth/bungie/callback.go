package bungie


import (
    "fmt"
    "errors"
    "net/http"
    "net/url"
    "io/ioutil"
    "strings"
    "strconv"
    "encoding/json"
    b64 "encoding/base64"
    "github.com/labstack/echo/v4"
    "gopkg.in/go-playground/validator.v9"
)

type (
    RespPayload struct {
        Code  string `json:"code" validate:"required"`
        State string `json:"state" validate:"required"`
    }

    ReqToken struct {
        GrantType string `json:"grant_type"`
        Code      string `json:"code"`
    }

    RespToken struct {
        AccessToken      string `json:"access_token"`
        TokenType        string `json:"token_type"`
        ExpiresIn        int    `json:"expires_in"`
        RefreshToken     string `json:"refresh_token"`
        RefreshExpiresIn int    `json:"refresh_expires_in"`
        MembershipID     int    `json:"membership_id"`
    }

    // Bungie.net's API is like LoW after you go against it you want to kill the person who designed it.
    ResponseMask struct {
        Response    *MembershipResponse `json:"Response" validate:"required"`
        Message     string              `json:"Message"`
        ErrorStatus string              `json:"ErrorStatus"`
        ErrorCode   string              `json:"ErrorCode"`
    }

    MembershipResponse struct {
        PrimaryMembershipId int   `json:"primaryMembershipId"`
        BungieNetUser       *User `json:"bungieNetUser" validate:"required"`
    }

    User struct {
        ID               string `json:"membershipId" validate:"required"`
        DisplayName      string `json:"displayName" validate:"required"`
        SteamDisplayName string `json:"steamDisplayname"`
        // TODO: PSN Display Name, need someone how has a PSN account
    }

    Response struct {
        Token     string `json:"token"`
        User      *User  `json:"user"`
    }
)

func getToken(p *RespPayload) (*RespToken, error) {
    client := new(http.Client)
    apiurl := fmt.Sprintf("%s/app/oauth/token", cfg.BaseAPI)
    authHeader := fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + secrets.Bungie)))

    data := url.Values{}
    data.Set("grant_type", "authorization_code")
    data.Set("code", p.Code)


    req, err := http.NewRequest("POST", apiurl, strings.NewReader(data.Encode()))
    if err != nil {
        log.Error(err)
        return nil, err
    }

    req.Header.Set("Authorization", authHeader)
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

    resp, err := client.Do(req)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    if resp.StatusCode != 200 {
        log.Debugln(string(body))
        err = errors.New(fmt.Sprintf("Return Code: %d", resp.StatusCode))
        log.Error(err)
        return nil, err
    }



    var payload RespToken
    json.Unmarshal(body, &payload)
    log.Debugln(payload)

    return &payload, nil
}

func GetUser(token  string) (*User, error) {
    client := new(http.Client)
    authHeader := fmt.Sprintf("%s %s", "Bearer", token)
    url := fmt.Sprintf("%s/User/GetMembershipsForCurrentUser", cfg.BaseAPI)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    req.Header.Set("Authorization", authHeader)
    req.Header.Set("X-API-Key", secrets.BungieAPIKey)

    resp, err := client.Do(req)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    if resp.StatusCode != 200 {
        err = errors.New(fmt.Sprintf("Server responded with code: %d", resp.StatusCode))
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return nil, err
    }

    var payload ResponseMask
    json.Unmarshal(body, &payload)

    log.Debugln(payload.Response.BungieNetUser)

    v := validator.New()
    if errs := v.Struct(payload); errs != nil {
        log.Error(errs)
        return nil, errs
    }

    return payload.Response.BungieNetUser, nil
}


func Callback(c echo.Context) (err error) {
    payload := new(RespPayload)
    if err = c.Bind(payload); err != nil {
        log.Error(err)
        return c.String(http.StatusBadRequest, "Error while processing payload")
    }

    v := validator.New()
    if errs := v.Struct(payload); errs != nil {
        return c.String(http.StatusBadRequest, "Invalid payload")
    }

    authPayload, err := getToken(payload)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Well rip Bungo did an upsie again")
    }

    accessToken := authPayload.AccessToken

    user, err := GetUser(accessToken)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Fuck me plz")
    }

    response := &Response{
        Token:     accessToken,
        User:      user,
    }

    return c.JSON(http.StatusOK, response)
}
