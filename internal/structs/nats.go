package structs

type (
    NATSRegistration struct {
        Id string `json:"id"`
    }

    NATS struct {
        SendRegistration chan *NATSRegistration
    }

    NATSConfig struct {
        URL string
    }
)
