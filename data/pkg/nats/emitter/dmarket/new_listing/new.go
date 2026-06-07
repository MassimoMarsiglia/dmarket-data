package newlisting

import "github.com/nats-io/nats.go"

type Config struct {
	Conn   *nats.Conn
	Lookup map[string]string
}

func New(cfg Config) *Emitter {
	return &Emitter{nc: cfg.Conn, lookup: cfg.Lookup}
}
