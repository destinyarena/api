package users

import(
    "context"
    "io"
    pb "github.com/arturoguerra/destinyarena-api/pkg/profiles"
    "google.golang.org/grpc"
    "github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
    conn, err := grpc.Dial(grpcfg.Profiles, grpc.WithInsecure())
    if err != nil {
        return c.String(502, err.Error())
    }

    defer conn.Close()

    p := pb.NewProfilesClient(conn)
    stream, err := p.GetAllProfiles(context.Background(), &pb.Empty{})
    if err != nil {
        return c.String(502, err.Error())
    }

    profiles := make([]*Profile, 0)
    for {
        profile, err := stream.Recv()
        if err == io.EOF {
            log.Infof("End of profiles stream")
            break
        }

        if err != nil {
            log.Error(err)
            return c.String(502, err.Error())
        }

        profiles = append(profiles, &Profile{
            Discord: profile.GetDiscord(),
            Bungie:  profile.GetBungie(),
            Faceit:  profile.GetFaceit(),
            Banned:  profile.GetBanned(),
        })
    }

    return c.JSON(200, profiles)
}
