package sales

import "github.com/nats-io/nats.go"

type Config struct {
	Conn *nats.Conn
}

func New(cfg Config) *Emitter {
	return &Emitter{nc: cfg.Conn}
}
