package structs

type (
    SQL struct {
        DBType   string `json:"dbtype"`
        Username string `json:"username"`
        Password string `json:"password"`
        Host     string `json:"host"`
        DBName   string `json:"name"`
    }

    Secrets struct {
        Discord      string `json:"discord"`
        Faceit       string `json:"faceit"`
        FaceitAPIKey string `json:"faceit_api_key"`
        Bungie       string `json:"bungie"`
        BungieAPIKey string `json:"bungie_api_key"`
        JWTSecret    string `json:"jwt_secret"`
        APIKey       string `json:"api_key"`
    }
)
