package faceit

import (
    "net/http"
)

type Faceit struct {
    UC *http.Client
    SC *http.Client
}

func New(token string) *Faceit {
    UC := newClient(token)
    SC := newClient(token)
    return &Faceit{
        UC: UC,
        SC: SC,
    }
}
