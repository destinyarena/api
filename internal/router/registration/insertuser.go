package registration

/*
This inserts new users into the database and returns errors if an alt tries to register
*/


import (
)

func insertUser(u *User) (err error, alt bool) {
    log.Infoln(u)
    err, _ = dbclient.RegisterUser(u.Discord, u.Bungie, u.Faceit)
    if err != nil {
        return err, true
    }

    return nil, false
}
