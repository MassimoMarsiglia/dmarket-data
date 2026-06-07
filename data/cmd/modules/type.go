package modules

import (
	"context"
)

type Module interface {
	Name() string
	Run(ctx context.Context) error
}
