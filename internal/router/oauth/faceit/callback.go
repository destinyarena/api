package faceit

import (
    "fmt"
    "bytes"
    "net/http"
    "io/ioutil"
    "encoding/json"
    b64 "encoding/base64"
    "github.com/labstack/echo/v4"
    "gopkg.in/go-playground/validator.v9"
)

type (
    ReqPayload struct {
        Code  string `json:"code" validate:"required"`
        State string `json:"state"`
    }

    ReqOAuthPayload struct {
        GrantType string `json:"grant_type"`
        Code      string `json:"code"`
    }

    RespOAuthPayload struct {
        AccessToken  string `json:"access_token"`
        RefreshToken string `json:"refresh_token"`
        ExpiresIn    string `json:"expires_in"`
        IdToken      string `json:"id_token"`
        Scope        string `json:"scope"`
    }

    User struct {
        PlayerID string `json:"player_id"`
        Nickname string `json:"nickname"`
    }
)

func getToken(p *ReqPayload) (RespOAuthPayload, error) {
   client := new(http.Client)
   OAuth2URL := fmt.Sprintf("%s/auth/v1/oauth/token", cfg.OAuthURL)
   authheader := b64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + secrets.Faceit))

   body := &ReqOAuthPayload{
       Code: p.Code,
       GrantType: "authorization_code",
   }

   bytesbody, err := json.Marshal(body)
   if err != nil {
       log.Error(err)
       return RespOAuthPayload{}, err
   }


   req, err := http.NewRequest("POST", OAuth2URL, bytes.NewBuffer(bytesbody))
   if err != nil {
       log.Error(err)
       return RespOAuthPayload{}, err
   }

   req.Header.Set("Authorization", authheader)
   req.Header.Set("Content-Type", "application/json")

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

   log.Infoln(string(authBody))

   var payload RespOAuthPayload
   json.Unmarshal(authBody, &payload)

   return payload, nil
}

func getProfile(p *RespOAuthPayload) (User, error) {
   client := new(http.Client)
   authheader := b64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + secrets.Faceit))
   userinfourl := fmt.Sprintf("%s/auth/v1/resources/userinfo")
   req, err := http.NewRequest("GET", userinfourl, nil)
   if err != nil {
       log.Error(err)
       return User{}, err
   }

   req.Header.Set("Authorization", authheader)

   resp, err := client.Do(req)
   if err != nil {
       log.Error(err)
       return User{}, err
   }

   defer resp.Body.Close()

   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
       log.Error(err)
       return User{}, err
   }

   var payload User
   json.Unmarshal(body, &payload)

   return payload, nil
}


func Callback(c echo.Context) error {
   payload := new(ReqPayload)
   if err := c.Bind(payload); err != nil {
       return c.String(http.StatusBadRequest, "Invalid or missing code")
   }

   v := validator.New()
   if errs := v.Struct(payload); errs != nil {
       return c.String(http.StatusBadRequest, "Invalid Payload")
   }


   authPayload, err := getToken(payload)
   if err != nil {
       return c.String(http.StatusInternalServerError, "Well rip it's not like faceit matters that much lol")
   }

   user, err := getProfile(&authPayload)
   if err != nil {
       return c.String(http.StatusInternalServerError, "Welp rip again :(")
   }

   return c.JSON(http.StatusOK, user)
}
