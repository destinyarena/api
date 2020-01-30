package config

import (
    "os"
)

func LoadHTTPConfig() (string, string) {
   log.Info("Loading HTTP Config")
   host := os.Getenv("HOST")
   if host == "" {
       host = "0.0.0.0"
   }

   port := os.Getenv("PORT")
   if port == "" {
       port = "8080"
   }

   return host, port
}
