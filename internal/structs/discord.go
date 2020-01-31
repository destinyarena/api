package structs

type (
    Discord struct {
        RemoteAPI   string `json:"remote_api"`
        ClientID    string `json:"client_id"`
        Scope       string `json:"scope"`
        RedirectURI string `json:"redirect_url"`
        BaseURL     string `json:"base_url"`
    }
)
