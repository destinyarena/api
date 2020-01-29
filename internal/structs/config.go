package structs

type (
    SQL struct {
        DBType  string `json:"dbtype"`
        Username string `json:"username"`
        Password string `json:"password"`
        Host     string `json:"host"`
    }

    HTTPServer struct {
        Port string `json:"port"`
        Host string `json:"host"`
    }

    Secrets struct {
        Discord   string `json:"discord"`
        Faceit    string `json:"faceit"`
        Bungie    string `json:"bungie"`
        JWTSecret string `json:"jwtsecret"`
        APIKey    string `json:"apikey"`
    }
)
