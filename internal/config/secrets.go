package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadSecrets() *structs.Secrets {
    log.Infoln("Loading Secrets")
    discord := os.Getenv("SECRET_DISCORD")
    faceit  := os.Getenv("SECRET_FACEIT")
    bungie  := os.Getenv("SECRET_BUNGIE")
    jwt     := os.Getenv("SECRET_JWT")
    apikey  := os.Getenv("SECRET_API_KEY")

    return &structs.Secrets{
        Discord:   discord,
        Faceit:    faceit,
        Bungie:    bungie,
        JWTSecret: jwt,
        APIKey:    apikey,
    }
}
