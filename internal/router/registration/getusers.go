package registration

import (
    discordAuth "github.com/arturoguerra/destinyarena-api/internal/router/oauth/discord"
    bungieAuth  "github.com/arturoguerra/destinyarena-api/internal/router/oauth/bungie"
    faceitAuth  "github.com/arturoguerra/destinyarena-api/internal/router/oauth/faceit"
)

func getUsers(p *Payload) (discord, faceit, bungie string, err  error) {
    rawDiscord, err := discordAuth.GetUser(p.Discord)
    if err != nil {
        log.Error(err)
        return "", "", "", err
    }

    rawFaceit, err := faceitAuth.GetProfile(p.Faceit)
    if err != nil {
        log.Error(err)
        return "", "", "", err
    }


    rawBungie, err := bungieAuth.GetUser(p.Bungie)
    if err != nil {
        log.Error(err)
        return "", "", "", err
    }

    return rawDiscord.ID, rawFaceit.PlayerID, rawBungie.ID, nil
}
