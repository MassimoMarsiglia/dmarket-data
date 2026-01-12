package pool

import (
	"context"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/subscription"
)

func (p *Pool[T, Z]) Start(ctx context.Context, startDelay time.Duration) error {
	subs := make([]subscription.Subscription[T], 0, len(p.workers))

	go func() {
		for _, worker := range p.workers {
			time.Sleep(startDelay)
			subs = append(subs, worker.Publisher.Subscribe())
			worker := worker // capture loop variable
			go func() {
				err := worker.Start(ctx)
				if err != nil {
					p.Publisher.NotifyErr(err)
				}
			}()
		}

		for _, sub := range subs {
			sub := sub
			go func() {
				defer sub.Close()

				for {
					select {
					case resp, ok := <-sub.C:
						if !ok {
							return
						}

						results, err := p.Transform(resp)
						if err != nil {
							p.Publisher.NotifyErr(err)
							continue
						}

						for _, result := range results {
							p.Publisher.Notify(result)
						}

					case err, ok := <-sub.E:
						if !ok {
							return
						}
						p.Publisher.NotifyErr(err)

					case <-ctx.Done():
						return
					}
				}
			}()
		}
	}()

	return nil
}
