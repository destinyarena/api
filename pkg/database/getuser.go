package database

/*
This Check if a user is already in the database and adds them if they are not
*/

import (
    "errors"
)

func (f *DBClient) GetUser(uid string) (error, *User) {
    db, err := f.Connect()
    if err != nil {
        return err, nil
    }

    defer db.Close()

    u := User{}
    if db.Where("discord = ? OR faceit = ? OR bungie = ?", uid, uid, uid).First(&u).RecordNotFound() {
        return errors.New("User not found"), nil
    }


    return nil, &u
}
