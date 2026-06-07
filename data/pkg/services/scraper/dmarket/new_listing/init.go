package newlisting

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
			cl := s.clients.PopFront()
			go func() {
				resp, err := s.getNewListings(s.context, cl)
				if err != nil {
					s.logger.Error("failed getting dmarket new listing", zap.Error(err))
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

		case <-s.context.Done():
			s.ticker.Stop()
			return nil
		}
	}
}
