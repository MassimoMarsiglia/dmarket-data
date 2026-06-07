package nats

import (
	"os"

	"github.com/nats-io/nats.go"
)

func Connect() (*nats.Conn, error) {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}
	return nats.Connect(url)
}
