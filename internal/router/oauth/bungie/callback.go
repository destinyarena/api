package bungie


import (
    "github.com/labstack/echo/v4"
)

type (
    RespPayload struct {
        Code  string `json:"code"`
        State string `json:"state"`
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

    User struct {
    }

    Response struct {
        Token     string `json:"token"`
        TokenType string `json:"token_type"`
        User      *User  `json:"user"`
    }
)

func getToken(p *RespPayload) (*RespToken, error) {
}

func GetUser(p *RespToken) (*User, error) {
}


func Callback(c echo.Context) error {
}
