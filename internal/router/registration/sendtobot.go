package registration

import (
    "github.com/arturoguerra/destinyarena-api/internal/structs"
)

func sendtoBot(u *User) {
    payload := &structs.NATSRegistration{
        Id: u.Discord,
    }

    log.Infoln("Sending payload to bot")
    nchan.SendRegistration <- payload
}
