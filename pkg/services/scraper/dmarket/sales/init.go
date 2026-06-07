package sales

import (
	"encoding/json"
	"log"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"go.uber.org/zap"
)

func (s *Service) init() error {

	reqChan := make(chan dmarket.AggregatorGetLastSalesParams)

	for {
		select {
		case req := <-reqChan:
			r := req
			s.priorityQueue.PushBack(&r)
		case <-s.ticker.C:
			if s.clients.Len() == 0 {
				log.Fatal(ErrNoClients)
			}

			if s.queue.Len() == 0 {
				log.Fatal(ErrNoItemsInQueue)
			}

			var req *dmarket.AggregatorGetLastSalesParams
			if s.priorityQueue.Len() == 0 {
				req = s.queue.PopFront()
				r := *req
				s.queue.PushBack(&r)
			} else {
				req = s.priorityQueue.PopFront()
			}

			cl := s.clients.PopFront()
			s.clients.PushBack(cl)
			go func() {
				resp, err := s.getSales(s.context, cl, req)
				if err != nil {
					s.logger.Error(
						"failed getting dmarket sales",
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

		case <-s.context.Done():
			s.ticker.Stop()
			close(reqChan)
			return nil
		}
	}
}
