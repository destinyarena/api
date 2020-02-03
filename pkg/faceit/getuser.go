package faceit

import (
    "fmt"
    "errors"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "gopkg.in/go-playground/validator.v9"
)

func (f *Faceit) GetUser(guid string) (*FaceitUser, error) {
    url := fmt.Sprintf("https://open.faceit.com/data/v4/players/%s", guid)

    req, _ := http.NewRequest("GET", url, nil)
    resp, err := f.SC.Do(req)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != 200 {
        err = errors.New("Invalid code")
        return nil, err
    }

    defer resp.Body.Close()

    rawbody, _ := ioutil.ReadAll(resp.Body)
    var body RawUser
    json.Unmarshal([]byte(rawbody), &body)

    v := validator.New()
    if err = v.Struct(body); err != nil {
        return nil, err
    }

    if _, ok := body.Games["destiny2"]; ok {
        user := &FaceitUser{
            Id:         body.Id,
            Username:   body.Username,
            SkillLevel: body.Games["destiny2"].SkillLevel,
            Steam:      body.SteamID,
        }

        return user, nil
    }

    return nil, errors.New("User doesn't have destiny 2 linked")
}
