package faceit

import (
    "fmt"
    "net/http"
)

type AddHeaderTransport struct {
   Token string
   T http.RoundTripper
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
   req.Header.Add("Content-Type", "application/json")
   req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", adt.Token))
   return adt.T.RoundTrip(req)
}

func newClient(token string) *http.Client {
    t := http.DefaultTransport
    return &http.Client{Transport: &AddHeaderTransport{
        Token: token,
        T: t,
    }}
}
