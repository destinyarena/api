package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadDiscordConfig() *structs.Discord {
    return &structs.Discord{
        ClientID:    os.Getenv("DISCORD_CLIENT_ID"),
        Scope:       os.Getenv("DISCORD_SCOPE"),
        RedirectURI: os.Getenv("DISCORD_REDIRECT_URI"),
        BaseURL:     os.Getenv("DISCORD_BASE_URL"),
    }
}
