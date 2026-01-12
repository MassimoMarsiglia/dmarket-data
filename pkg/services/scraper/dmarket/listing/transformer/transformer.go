package transformer

import (
	"encoding/json"
	"io"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client"
	list "github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/dmarket/listing/pool"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/dmarket/listing/transformer/cache"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool"
)

// Transform implements [pool.Transformer].
func (d *DmarketListingTransformer) Transform(r io.Reader, w io.Writer) error {
	// Read the entire response
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	var resp client.EntityGetItemsResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	for _, item := range resp.Objects {
		prevPrice := item.Price.USD
		cachedPrice, ok := d.cache.Get(item.ItemId)
		if ok && cachedPrice == prevPrice {
			continue
		}
		d.cache.Set(item.ItemId, item.Price.USD, 5*time.Minute)

		listing := list.ListingItem{
			Name:      item.Title,
			ItemID:    item.ItemId,
			Price:     item.Price.USD,
			Float:     item.Extra.FloatValue,
			OfferId:   item.Extra.OfferId,
			Tradable:  item.Extra.Tradable,
			PaintSeed: item.Extra.PaintSeed,
			// PaintSeed: item.,
		}
		encoder := json.NewEncoder(w)
		err = encoder.Encode(listing)
		if err != nil {
			return err
		}

	}
	return nil
}

type DmarketListingTransformer struct {
	cache       cache.Cache[string]
	CacheExpiry time.Duration
}

type DmarketListingTransformerConfig struct {
	CacheExpiry time.Duration
}

func New(cfg DmarketListingTransformerConfig) *DmarketListingTransformer {
	return &DmarketListingTransformer{
		CacheExpiry: cfg.CacheExpiry,
		cache:       cache.NewCache[string](),
	}
}

var _ pool.Transformer = (*DmarketListingTransformer)(nil)
