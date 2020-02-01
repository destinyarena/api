package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadBotConfig() *structs.Bot {
    return &structs.Bot{
        API: os.Getenv("BOT_API"),
    }
}
