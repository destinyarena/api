package structs

type Secrets struct {
    Discord       string `json:"discord"`
    Faceit        string `json:"faceit"`
    FaceitAPIKey  string `json:"faceit_api_key"`
    FaceitUserKey string `json:"faceit_user_key"`
    Bungie        string `json:"bungie"`
    BungieAPIKey  string `json:"bungie_api_key"`
    JWTSecret     string `json:"jwt_secret"`
    APIKey        string `json:"api_key"`
    DBPassword    string `json:"db_password"`
}
