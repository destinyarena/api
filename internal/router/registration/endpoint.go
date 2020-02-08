package registration

import (
    "errors"
    "net/http"
    "github.com/labstack/echo/v4"
    "gopkg.in/go-playground/validator.v9"
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
        Bungie    string `json:"bungie"    validate:"required"`
        Recaptcha string `json:"recaptcha" validate:"required"`
    }

    User struct {
        Discord string
        Bungie  string
        Faceit  string
    }
)

func endpoint(c echo.Context) (err error) {
    payload := new(Payload)
    if err = c.Bind(payload); err != nil {
        log.Error(err)
        return c.String(http.StatusBadRequest, "Invalid payload")
    }

    v := validator.New()
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
        return c.String(http.StatusUnauthorized, "Error fetching user profile from payload")
    }

    user := &User{
        Discord: discord,
        Bungie:  bungie,
        Faceit:  faceit,
    }

    log.Debugln(user)

    err, alt := insertUser(user)
    if alt {
        err = errors.New("Sorry but one or more account's are already registered")
        log.Error(err)
        return c.String(http.StatusForbidden, "Sorry but one or more accounts's are already registered")
    }

    if err != nil {
        log.Error(err)
        return c.String(http.StatusInternalServerError, "Well something went wrong while adding a new user")
    }

    sendtoBot(user)

    return c.String(http.StatusOK, "You have successfully registered")
}
