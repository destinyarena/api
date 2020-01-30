package structs

type (
    Faceit struct {
        ClientID string `json:"client_id"`
        OAuthURL string `json:"oauth_url"`
        BaseAPI  string `json:"base_api"`
    }
)
