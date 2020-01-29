package faceit

import (
    "fmt"
    "net/http"

var client *http.Client

type AddHeaderTransport struct {
   Token string
   T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
   req.Header.Add("User-Agent", "go")
   req.Header.Add("Content-Type", "application/json")
   req.Header.Add("Authorixation", fmt.Sprinf("Bearer %s", adt.Token)
   return adt.T.RoundTrip(req)
}

func NewClient(token string) *http.Client {
    t := http.DefaultTransport
    return &http.Client{Transport: &AddHeaderTransport{t}}
}


type Faceit {
    UC *http.Client
}

func New(token string) *Faceit {
    UC := NewClient(token)
    return &Faceit{UC}
}
