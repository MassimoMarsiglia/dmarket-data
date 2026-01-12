package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool/worker"
	"github.com/google/uuid"
)

type DmarketWorkerConfig struct {
	ID       uuid.UUID
	Dmarket  client.DmarketCfg
	ReqDelay time.Duration
}

type NewListingWorker struct {
	w *worker.Worker[*client.EntityGetItemsResponse]
}

func New(cfg DmarketWorkerConfig) (*NewListingWorker, error) {
	cl, err := client.NewDmarketClient(cfg.Dmarket)
	if err != nil {
		return nil, err
	}

	return &NewListingWorker{
		w: worker.New(worker.Config[*client.EntityGetItemsResponse]{
			ID:     cfg.ID,
			Logger: nil,
			Work: func(ctx context.Context) (*client.EntityGetItemsResponse, error) {
				resp, err := cl.GetMarketItemsWithResponse(ctx, &client.GetMarketItemsParams{
					GameId:   "a8db",
					Limit:    utils.Ptr(20),
					OrderBy:  utils.Ptr("updated"),
					Currency: "USD",
				})
				if err != nil {
					return nil, err
				}
				code := resp.StatusCode()
				if resp.StatusCode() != 200 {
					return nil, fmt.Errorf("bad response %s %d", resp.Status(), code)
				}

				resp200 := resp.JSON200
				return resp200, nil
			},
			Delay: cfg.ReqDelay,
		}),
	}, nil
}
