package registration

/*
This inserts new users into the database and returns errors if an alt tries to register
*/


import (
)

func insertUser(u *User) (err error, alt bool) {
    err, user := dbclient.RegisterUser(u.Discord, u.Bungie, u.Faceit)
    if err != nil && user == nil {
        log.Error(err)
        return err, false
    }

    if err != nil && user != nil {
        log.Error(err)
        return err, true
    }

    return nil, false
}
