package structs

type (
    Bungie struct {
        ClientID  string `json:"client_id"`
        OAuth2URL string `json:"oauth2_url"`
        BaseAPI   string `json:"base_api"`
        Scope     string `json:"scope"`
    }
)
