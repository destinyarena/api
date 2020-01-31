package registration

import (
    "fmt"
    "net/http"
    "bytes"
    "encoding/json"
    "io/ioutil"
    "errors"
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
        err = errors.New(fmt.Sprintf("Returned Code: %d", resp.StatusCode))
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
