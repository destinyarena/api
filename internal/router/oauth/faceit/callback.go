package faceit

import (
    "fmt"
    //"bytes"
    "time"
    "errors"
    "net/http"
    "io/ioutil"
    "encoding/json"
    b64 "encoding/base64"
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
    "gopkg.in/go-playground/validator.v9"
    "github.com/arturoguerra/destinyarena-api/internal/utils"
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
        TokenType    string `json:"token_type"`
        RefreshToken string `json:"refresh_token"`
        ExpiresIn    string `json:"expires_in"`
        IdToken      string `json:"id_token"`
        Scope        string `json:"scope"`
    }

    User struct {
        GUID     string `json:"guid"`
        Nickname string `json:"nickname"`
    }

    Claims struct {
        *User
        jwt.StandardClaims
    }
)

func getToken(p *ReqPayload) (*RespOAuthPayload, error) {
   client := new(http.Client)

   authheader := fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + secrets.Faceit)))
   authurl := fmt.Sprintf("%s/auth/v1/oauth/token?grant_type=authorization_code&code=%s", cfg.BaseAPI, p.Code)


   req, err := http.NewRequest("POST", authurl, nil)
   if err != nil {
       log.Error(err)
       return nil, err
   }

   log.Debugln(authheader)
   log.Debugln(authurl)

   req.Header.Set("Authorization", authheader)
   req.Header.Set("Content-Type", "application/json")

   resp, err := client.Do(req)
   if err != nil {
       log.Error(err)
       return nil, err
   }

   if resp.StatusCode != 200 {
       err = errors.New(fmt.Sprintf("Server didn't Respond with a 200"))
       log.Error(err)
       return nil, err
   }

   defer resp.Body.Close()

   authBody, err := ioutil.ReadAll(resp.Body)
   if err != nil {
       log.Error(err)
       return nil, err
   }

   var payload RespOAuthPayload
   json.Unmarshal(authBody, &payload)
   log.Debugln(payload)

   return &payload, nil
}

func GetProfile(token  string) (*User, error) {
   client := new(http.Client)

   authheader := fmt.Sprintf("Bearer %s", token)
   userinfourl := fmt.Sprintf("%s/auth/v1/resources/userinfo", cfg.BaseAPI)
   req, err := http.NewRequest("GET", userinfourl, nil)
   if err != nil {
       log.Error(err)
       return nil, err
   }

   req.Header.Set("Authorization", authheader)

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
       log.Error(string(body))
       err := errors.New(fmt.Sprintf("Server responded with error code: %d", resp.StatusCode))
       return nil, err
   }

   var payload User
   json.Unmarshal(body, &payload)

   return &payload, nil
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

   accessToken := authPayload.AccessToken


   user, err := GetProfile(accessToken)
   if err != nil {
       return c.String(http.StatusInternalServerError, "Welp rip again :(")
   }

    claims := &Claims{
        User: user,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt:  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
            IssuedAt: time.Now().Unix(),
        },
    }

   token, err := utils.SignJWT(claims)
   if err != nil {
       log.Error(err)
       return c.String(http.StatusInternalServerError, "Error while creating Token")
   }

   r := map[string]interface{}{
       "token": token,
   }

   return c.JSON(http.StatusOK, r)
}
