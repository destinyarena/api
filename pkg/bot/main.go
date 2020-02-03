package bot

import (
    "net/http"
)

type Bot struct {
    Base    string
    Client *http.Client
}


func New(token, base string) *Bot {
    client := newClient(token)
    return &Bot{
        client,
        base,
    }
}
