package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadSecrets() *structs.Secrets {
    discord       := os.Getenv("SECRET_DISCORD")
    faceit        := os.Getenv("SECRET_FACEIT")
    faceitapikey  := os.Getenv("SECRET_FACEIT_API_KEY")
    faceituserkey := os.Getenv("SECRET_FACEIT_USER_KEY")
    bungie        := os.Getenv("SECRET_BUNGIE")
    bungieapikey  := os.Getenv("SECRET_BUNGIE_API_KEY")
    jwt           := os.Getenv("SECRET_JWT")
    apikey        := os.Getenv("SECRET_API_KEY")
    dbpassword    := os.Getenv("SECRET_DB_PASSWORD")

    return &structs.Secrets{
        Discord:       discord,
        Faceit:        faceit,
        FaceitAPIKey:  faceitapikey,
        FaceitUserKey: faceituserkey,
        Bungie:        bungie,
        BungieAPIKey:  bungieapikey,
        JWTSecret:     jwt,
        APIKey:        apikey,
        DBPassword:    dbpassword,
    }
}
