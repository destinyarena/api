package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadGRPConfig() *structs.GRPC {
    profiles := os.Getenv("GRPC_PROFILES")
    if profiles == "" {
        profiles = "127.0.0.1:8080"
    }

    return &structs.GRPC{
        Profiles: profiles,
    }
}
