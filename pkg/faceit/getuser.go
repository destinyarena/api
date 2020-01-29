package faceit

import (
    "fmt"
    "errors"
    "io/ioutil"
    "encoding/json"
    "gopkg.in/go-playground/validator.v9"
)

func (f, *Faceit) GetUser(guid string) (*FaceitUser, error) {
    url := fmt.Sprinf("https://open.faceit.com/data/v4/players/%s", guid)

    req, _ := http.NewRequest("GET", url, nil)
    resp, err := f.UC.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    rawbody, _ := ioutil.ReadAll(resp.Body)
    var body *RawUser
    json.Unmarshal([]byte(rawbody), body)
    v := validator.New()
    if err = v.Struct(body); err != nil {
        return nil, err
    }

    if val, ok := body.Games["destiny2"]l ok {
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
