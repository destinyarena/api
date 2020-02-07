package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadGRPConfig() *structs.GRPC {
    phost := os.Getenv("GRPC_PROFILES_HOST")
    if phost == "" {
        phost = "127.0.0.1"
    }

    pport := os.Getenv("GRPC_PROFILES_PORT")
    if pport == "" {
        pport = "8080"
    }

    return &structs.GRPC{
        ProfilesHost: phost,
        ProfilesPort: pport,
    }
}
