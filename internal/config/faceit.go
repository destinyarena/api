package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadFaceitConfig() *structs.Faceit {
    return &structs.Faceit{
        ClientID: os.Getenv("FACEIT_CLIENT_ID"),
        OAuthURL: os.Getenv("FACEIT_OAUTH_URL"),
        BaseAPI:  os.Getenv("FACEIT_BASE_API"),
    }
}
