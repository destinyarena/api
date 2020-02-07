package profiles

import (
    "fmt"
    "google.golang.org/grpc"
)

func New(host, port string) (ProfilesClient, error) {
    address := fmt.Sprintf("%s:%s", host, port)
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, err
    }

    c := NewProfilesClient(conn)

    return  c, nil
}
