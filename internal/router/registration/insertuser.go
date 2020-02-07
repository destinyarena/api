package registration

/*
This inserts new users into the database and returns errors if an alt tries to register

TODO: Fix this whole file by moving to it to its own package called botapi
*/


import (
    "fmt"
    "net/http"
    "context"
    "github.com/arturoguerra/destinyarena-api/pkg/profiles"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    // "github.com/arturoguerra/destinyarena-api/pkg/bot"
)

var botcfg = config.LoadBotConfig()
var secrets = config.LoadSecrets()

type BotPayload struct {
    Discord string `json:"discord"`
    Skillvl int    `json:"skillvl"`
    Faceit  string `json:"faceit"`
}

func postToBot(uid string) error {
    url := fmt.Sprintf("%s/roles/%s", botcfg.API, uid)

    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        return err
    }

    req.Header.Set("Authorization", "Bearer " + secrets.APIKey)
    req.Header.Set("Content-Type", "application/json")

    client := new(http.Client)
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    if resp.StatusCode != 200 {
        err = fmt.Errorf("Server returned: %d", resp.StatusCode)
        return err
    }

    return nil
}

func insertUser(u *User) (err error, alt bool) {
    log.Debugln(u)
    _ , err = uClient.CreateProfile(context.Background(), &profiles.ProfileRequest{
        Discord: u.Discord,
        Bungie: u.Bungie,
        Faceit: u.Faceit,
    })
    if err != nil {
        return err, true
    }

    if err = postToBot(u.Discord); err != nil {
        log.Error(err)
        return err, false
    }

    log.Infof("User: %s has registered", u.Discord)

    return nil, false
}
