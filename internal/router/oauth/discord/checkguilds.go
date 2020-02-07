package discord

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type (
    Guild struct {
        Id string   `json:"id"`
        Name string `json:"name"`
    }
)

func inGuild(guilds []Guild) bool {
    gid := cfg.GuildID
    log.Infof("Arena Guild: %s", gid)
    for _, guild := range guilds {
        log.Infof("%s --- %s", guild.Id, guild.Name)
        if guild.Id == gid  {
            return true
        }
    }

    return false
}

func checkGuilds(token string) (error, bool) {
    client := new(http.Client)
    token = fmt.Sprintf("%s %s", "Bearer", token)
    log.Debug(token)

    req, err := http.NewRequest("GET", fmt.Sprintf("%s/v6/users/@me/guilds", cfg.BaseURL), nil)
    if err != nil {
        log.Error(err)
        return err, false
    }

    req.Header.Set("Authorization", token)

    resp, err := client.Do(req)
    if err != nil {
        log.Error(err)
        return err, false
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err, false
    }

    if resp.StatusCode != 200 && resp.StatusCode != 201 {
        log.Error(string(body))
        return fmt.Errorf("Server returned: %d", resp.StatusCode), false
    }

    guilds := make([]Guild, 0)
    json.Unmarshal(body, &guilds)

    return nil, inGuild(guilds)
}
