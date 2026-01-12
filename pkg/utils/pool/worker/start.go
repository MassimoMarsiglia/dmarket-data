package worker

import (
	"context"
	"time"
)

func (w *Worker[T]) Start(ctx context.Context) error {
	t := time.NewTicker(w.delay)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			w.Work(ctx)
		}
	}
}
