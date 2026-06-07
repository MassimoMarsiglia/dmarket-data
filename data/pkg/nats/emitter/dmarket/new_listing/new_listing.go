package newlisting

import "github.com/nats-io/nats.go"

type Emitter struct {
	nc     *nats.Conn
	lookup map[string]string
}
