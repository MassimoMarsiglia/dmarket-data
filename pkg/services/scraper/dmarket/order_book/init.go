package order_book

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
			req := s.queue.PopFront()
			go func() {
				resp, err := s.getOrderBook(s.context, cl, req)
				if err != nil {
					s.logger.Error(
						"failed getting dmarket order book",
						zap.Any("request", req),
						zap.Error(err),
					)
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
			s.queue.PushBack(req)

		case <-s.context.Done():
			s.ticker.Stop()
			return nil
		}
	}
}
