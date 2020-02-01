package faceit

type (
    FaceitUser struct {
        Id         string `json:"id"`
        Username   string `json:"username"`
        SkillLevel int    `json:"skilllevel"`
        Steam      string `json:"steam"`
    }

    Game struct {
        SkillLevel int `json:"skill_level" validate:"required"`
    }

    RawUser struct {
        Id       string `json:"player_id" validate:"required"`
        Username string `json:"nickname" validate:"required"`
        SteamID  string `json:"steam_id_64" validate:"required"`
        Games    map[string]Game `json:"games" validate:"required"`
    }
)
