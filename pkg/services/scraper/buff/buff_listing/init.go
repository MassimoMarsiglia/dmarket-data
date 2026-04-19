package bufflisting

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
)

func (s *Service) init() error {
	for {
		select {
		case <-s.ticker.C:
			if s.clients.Len() == 0 {
				log.Fatal(ErrNoClients)
			}

			if s.queue.Len() == 0 {
				log.Fatal(ErrNoItemsInQueue)
			}

			cl := s.clients.PopFront()
			item := s.queue.PopFront()
			go func() {

				resp, err := s.getListings(s.context, cl, *item)
				if err != nil {
					s.logger.Error("failed getting listing", zap.Error(err))
					return
				}

				for _, item := range resp {
					itemb, err := json.Marshal(item)
					if err != nil {
						return
					}
					s.nc.Publish(NATS_KEY, itemb)
				}

			}()
			s.clients.PushBack(cl)
			s.queue.PushBack(item)

		case <-s.context.Done():
			s.ticker.Stop()
			return nil
		}
	}
}
