package registration

import (
    "github.com/arturoguerra/destinyarena-api/internal/router/oauth/discord"
    "github.com/arturoguerra/destinyarena-api/internal/router/oauth/bungie"
    "github.com/arturoguerra/destinyarena-api/internal/router/oauth/faceit"
    "github.com/arturoguerra/destinyarena-api/internal/utils"
    "gopkg.in/go-playground/validator.v9"
)

func getUsers(p *Payload) (string, string, string, error) {
    v := validator.New()

    discordClaims := new(discord.Claims)
    err := utils.DecryptJWT(p.Discord, discordClaims)
    if err != nil {
        log.Error(err)
        return "", "", "", err
    }

    faceitClaims := new(faceit.Claims)
    err = utils.DecryptJWT(p.Faceit, faceitClaims)
    if err != nil {
        log.Error(err)
        return "", "", "", err
    }


    bungieClaims := new(bungie.Claims)
    err = utils.DecryptJWT(p.Bungie, bungieClaims)
    if err != nil {
        log.Error(err)
        return "", "", "", err
    }

    if err = v.Struct(discordClaims.User); err != nil {
        log.Error(err)
        return "", "", "", err
    }
    if err = v.Struct(faceitClaims.User); err != nil {
        log.Error(err)
        return "", "", "", err
    }

    if err = v.Struct(bungieClaims.User); err != nil {
        log.Error(err)
        return "", "", "", err
    }


    return discordClaims.ID, faceitClaims.GUID, bungieClaims.ID, nil
}
