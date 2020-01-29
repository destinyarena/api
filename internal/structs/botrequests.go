package main

type (
    // Bot response to get USER by ID
    BotGetUser struct {
        ID             int    `json:"id"`
        Username       string `json:"username"`
        Discriminator  string `json:"discriminator"`
        Admin          bool   `json:"admin"`
    }
)
