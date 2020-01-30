package structs

type (
    // POST /api/registration/submit
    ReqRegistrationSubmit struct {
        Discord   string `json:"discord"`
        Faceit    string `json:"faceit"`
        Bungie    string `json:"bungie"`
        Recaptcha string `json:"recaptcha"`
    }

    // GET /api/users/get/{discord,faceit,bungie}/{user_id}
    RespGetUserByID struct {
        Discord  *DiscordUser `json:"discord"`
        Faceit   *FaceitUser  `json:"faceit"`
        Bungie   *BungieUser  `json:"bungie"`
        Banned   bool         `json:"banned"`
    }

    // POST /api/users/ban
    ReqUserBan struct {
        Discord string `json:"discord"`
        Faceit  string `json:"faceit"`
        Bungie  string `json:"bungie"`
    }

    // POST /api/users/unban
    ReqUserUnban struct {
        *ReqUserBan
    }

    // GET /api/users/list
    RespUserList struct {
        Count  int `json:"count"`
        Users  []*SimpleUser `json:"users"`
    }

    DiscordUser struct {
        Id int `json:"id"`
        Username string `json:"username"`
        Discriminator string `json:"discriminator"`
    }

    FaceitUser struct {
        Id string       `json:"id"`
        Username string `json:"username"`
        SkillLevel int  `json:"skilllevel"`
        Steam string    `json:"steam"`
    }

    BungieUser struct {
        Id string `json:"id"`
        Username string `json:"username"`
    }

    SimpleUser struct {
        Discord string `json:"discord"`
        Faceit  string `json:"faceit"`
        Bungie  string `json:"bungie"`
        Banned  bool   `jsoN:"banned"`
    }
)
