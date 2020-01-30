package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadBungieConfig() *structs.Bungie {
    return &structs.Bungie{
        ClientID:  os.Getenv("BUNGIE_CLIENT_ID"),
        OAuth2URL: os.Getenv("BUNGIE_OAUTH_URL"),
        BaseAPI:   os.Getenv("BUNGIE_BASE_API"),
        Scope:     os.Getenv("BUNGIE_SCOPE"),
    }
}
