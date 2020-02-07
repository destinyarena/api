package faceit

import (
    "fmt"
    "google.golang.org/grpc"
)

func New(host, port string) (FaceitClient, error) {
    address := fmt.Sprintf("%s:%s", host, port)
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, err
    }

    c := NewFaceitClient(conn)

    return  c, nil
}
