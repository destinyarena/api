package config

import (
    "os"
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func LoadSQLConfig() *structs.SQL {
    log.Infoln("Loading SQL Config")
    dbtype := os.Getenv("DB_TYPE")
    username := os.Getenv("DB_USERNAME")
    host := os.Getenv("DB_HOST")
    dbname := os.Getenv("DB_NAME")

    return &structs.SQL{
        DBType:   dbtype,
        Username: username,
        Host:     host,
        DBName:   dbname,
    }
}
