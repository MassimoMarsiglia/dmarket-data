package newlisting

import (
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
					if err := s.em.Emit(item); err != nil {
						s.logger.Error("failed emitting new listing", zap.Error(err))
					}
				}

			}()
			s.clients.PushBack(cl)

		case <-s.context.Done():
			s.ticker.Stop()
			return nil
		}
	}
}
