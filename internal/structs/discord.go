package structs

type (
    Discord struct {
        ClientID    string `json:"client_id"`
        Scope       string `json:"scope"`
        RedirectURI string `json:"redirect_url"`
        BaseURL     string `json:"base_url"`
    }
)
