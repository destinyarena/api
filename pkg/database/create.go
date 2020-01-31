package database

/*
This Check if a user is already in the database and adds them if they are not
*/

import (
    "errors"
)

func (f *DBClient) RegisterUser(discord, bungie, faceit string) (error, *User) {
    db, err := f.Connect()
    if err != nil {
        return err, nil
    }

    defer db.Close()

    u := User{}
    if db.Where("discord = ? OR faceit = ? OR bungie = ?", discord, faceit, bungie).First(&u).RecordNotFound() {
        newUser := &User{
            Discord: discord,
            Bungie:  bungie,
            Faceit:  faceit,
        }

        db.Create(newUser)

        return nil, newUser
    }

    err = errors.New("Member already exists")

    return err, &u
}
