package listing

import (
	"sync"
	"testing"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/dmarket/listing/pool"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/dmarket/listing/transformer"
	poolType "github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool"
	"github.com/stretchr/testify/assert"
)

func Test_Listings(t *testing.T) {
	transformer := transformer.New(transformer.DmarketListingTransformerConfig{
		CacheExpiry: 10 * time.Minute,
	})
	pool := pool.New(pool.DmarketPoolConfig{
		ReqDelay: 2 * time.Second,
		Transformers: []poolType.TransformFunc{
			transformer.Transform,
		},
	})
	cfg := ServiceCfg{
		Ctx:        t.Context(),
		Pool:       pool,
		StartDelay: 1 * time.Second,
		NumWorkers: 5,
	}
	svc := New(cfg)
	sub := svc.publisher.Subscribe()
	err := svc.StartFeed()
	assert.NoError(t, err)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			select {
			case item, ok := <-sub.C:
				if !ok {
					t.Log("Item not okay")
					return
				}
				t.Logf("Received item: %+v", item)
			case err, ok := <-sub.E:
				if !ok {
					return
				}
				t.Errorf("Received error: %+v", err)
			}
		}
	}()
	wg.Wait()
}
