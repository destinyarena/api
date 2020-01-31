package utils

import (
    "net/url"
    "crypto/rand"
    "encoding/hex"
)


// Makes a Query Param safe url
func SafeUrl(unsafe string) string {
    return url.QueryEscape(unsafe)
}

// Created a hexadecimal string
func RandomHex(n int) (string, error) {
  bytes := make([]byte, n)
  if _, err := rand.Read(bytes); err != nil {
    return "", err
  }
  return hex.EncodeToString(bytes), nil
}
