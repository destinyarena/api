package faceit

type (
    // GetUser
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
        Games    map[string]Game `json:"games" validate:"required"`
    }


    //GetInvite
    ReqInvitePayload struct {
        EntityID   string `json:"entity_id"`
        EntityType string `json:"entity_type"`
        Type       string `json:"type"`
        MaxAge     int    `json:"max_age"`
        MaxUsers   int    `json:"max_users"`
    }

    Invite struct {
        Payload InvitePayload `json:"payload"`
    }

    InvitePayload struct {
        Code string `json:"code"`
    }
)
