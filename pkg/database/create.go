

func (f *Faceit) RegisterUser(NewUser *User) error {
    err, db := f.Connect()
    if err != nil {
        return err
    }

    u := new(User)
    db.Find(u, "discord <> ? OR faceit <> ? OR bungie <> ?", NewUser.Discord, NewUser.Faceit, NewUser.Bungie)
    if u.ID != "" {
        return errors.New("User already exists")
    }

    db.Create(NewUser)
    return nil
}
