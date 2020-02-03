package structs

type (
    SQL struct {
        DBType   string `json:"dbtype"`
        Username string `json:"username"`
        Host     string `json:"host"`
        DBName   string `json:"name"`
    }
)
