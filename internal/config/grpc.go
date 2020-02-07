package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadGRPConfig() *structs.GRPC {
    return &structs.GRPC{
        ProfilesHost: os.Getenv("GRPC_PROFILES_HOST"),
        ProfilesPort: os.Getenv("GRPC_PROFILES_PORT"),
    }
}
