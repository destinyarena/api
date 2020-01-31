package registration

import (
    "github.com/labstack/echo/v4"
)


type (
    RecaptchaPayload struct {
        Secret string `json:"secret"`
    }

    ResponseRecaptcha struct {
        Success bool `json:"bool"`
    }

    Payload struct {
        Discord   string `json:"discord"   validate:"required"`
        Faceit    string `json:"faceit"    validate:"required"`
        Bungie    string `json:"bungie     validate:"required"`
        Recaptcha string `json:"recaptcha" validate:"required"`
    }
)

func recaptcha(code string) (err error) {
    client := new(http.Client)
    rurl := "https://www.google.com/recaptcha/api/siteverify"
    body, err := json.Marshal(&RecaptchaPayload{
        Secret: code,
    })

    if err != nil {
        log.Error(err)
        return err
    }


    payload := bytes.NewBuffer(body)

    req, err := http.NewRequest("POST", rurl, payload)
    if err != nil {
        log.Error(err)
        return err
    }

    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if err != nil {
        log.Error(err)
        return err
    }

    if resp.StatusCode != 200 {
        err = errors.New("Returned Code: %d", resp.StatusCode)
        log.Error(err)
        return err
    }

    defer resp.Body.Close()

    bodyresp, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error(err)
        return err
    }

    var p ResponseRecaptcha
    json.Unmarshal(bodyresp, &p)

    if !p.Success {
        err = errors.New("Looks like the token expired")
        log.Error(err)
        return err
    }

    return nil
}


func endpoint(c echo.Context) (err error) {
    payload := new(payload)
    if err = c.Bind(payload); err != nil {
        log.Error(err)
        return c.String(http.StatusBadRequest, "Invalid payload")
    }

    v := validate.New()
    if err = v.Struct(payload); err != nil {
        log.Error(err)
        return c.String(http.StatusBadRequest, "Error validating payload")
    }

//    if err = recaptcha(payload.Recaptcha); err != nil {
//        log.Error(err)
//        return c.String(404, "Invalid Recaptcha")
//    }

    discord, bungie, faceit, err := getUsers(payload)
    if err != nil {
        log.Error(err)
        return c.String(http.BadRequest, "Error fetching user profile from payload")
    }

    user := newUser(discord, bungie, faceit)

    if err, alt := insertUser(user)
    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Well something went wrong while adding a new user")
    }

    if alt {
        err = errors.New("Sorry but one or more account's are already registered")
        log.Error(err)
        return c.String(401, "Sorry but one or more accounts's are already registered")
    }

    return c.String(http.StatusOK, "You have successfully registered")
}
