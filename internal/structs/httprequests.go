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
    RespGetUserByID {
        Discord  *DiscordUser `json:"discord"`
        Faceit   *FaceitUser  `json:"faceit"`
        Bungie   *BungieUser  `json:"bungie"`
        Banned   bool         `json:"banned"`
    }

    // POST /api/users/ban
    ReqUserBan {
        Discord string `json:"discord"`
        Faceit  string `json:"faceit"`
        Bungie  string `json:"bungie"`
    }

    // POST /api/users/unban
    ReqUserUnban {
        *ReqUserBan
    }

    // GET /api/users/list
    RespUserList {
        Count  int `json:"count"`
        Users  []*SimpleUser `json:"users"`
    }

    SimpleUser {
        Discord string `json:"discord"`
        Faceit  string `json:"faceit"`
        Bungie  string `json:"bungie"`
        Banned  bool   `jsoN:"banned"`
    }
)
