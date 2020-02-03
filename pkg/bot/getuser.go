package bot

type DUser struct {
    ID       string   `json:"id"       validate:"required"`
    Username string   `json:"username" validate:"required"`
    Roles    []string `json:"roles"    validate:"required"`
    Admin    bool     `json:"admin"    validate:"required"`
}
