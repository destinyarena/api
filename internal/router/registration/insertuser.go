package registration

import (
    "fmt"
    "context"
    "google.golang.org/grpc"
    "github.com/arturoguerra/destinyarena-api/pkg/profiles"
    "github.com/arturoguerra/destinyarena-api/internal/config"
    pb "github.com/arturoguerra/destinyarena-api/pkg/profiles"
)

var botcfg = config.LoadBotConfig()
var secrets = config.LoadSecrets()

func insertUser(u *User) (err error, alt bool) {
    grpcfg := config.LoadGRPConfig()
    addr := fmt.Sprintf("%s:%s", grpcfg.ProfilesHost, grpcfg.ProfilesPort)
    conn, err := grpc.Dial(addr, grpc.WithInsecure())
    if err != nil {
        return err, false
    }

    defer conn.Close()

    client := pb.NewProfilesClient(conn)
    _ , err = client.CreateProfile(context.Background(), &profiles.ProfileRequest{
        Discord: u.Discord,
        Bungie: u.Bungie,
        Faceit: u.Faceit,
    })
    if err != nil {
        return err, true
    }

    log.Infof("User: %s has registered", u.Discord)

    return nil, false
}
