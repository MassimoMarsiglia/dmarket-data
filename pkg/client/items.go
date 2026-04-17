package client

import (
	"strconv"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

type FilterFunc[T any] func(T) (bool, error)

func PriceIDFilter() FilterFunc[models.Item] {
	mu := sync.Mutex{}
	seenItems := make(map[string]int64)

	return func(i models.Item) (bool, error) {
		mu.Lock()
		defer mu.Unlock()

		prev, ok := seenItems[i.ItemID]
		if ok && prev == i.Price {
			return false, nil
		}

		seenItems[i.ItemID] = i.Price
		return true, nil
	}
}

func (e EntityGetItemsResponse) Items(filters ...FilterFunc[models.Item]) ([]models.Item, error) {
	items := make([]models.Item, 0, len(e.Objects))

Transform:
	for _, item := range e.Objects {

		price, err := strconv.ParseInt(item.Price.USD, 10, 32)
		if err != nil {
			return nil, err
		}
		owner := models.Owner{
			ID:     item.OwnerDetails.Id,
			Wallet: item.OwnerDetails.Wallet,
		}

		var stickers []models.Sticker
		if item.Extra.Stickers != nil {
			stickers = make([]models.Sticker, 0, len(*item.Extra.Stickers))
			for _, sticker := range *item.Extra.Stickers {
				stickers = append(stickers, models.Sticker{
					Name:  sticker.Name,
					Image: sticker.Name,
				})
			}
		}
		i := models.Item{
			MarketHashName: item.Title,
			Image:          item.Image,
			Price:          price,
			Owner:          owner,
			Float:          item.Extra.FloatValue,
			PaintSeed:      item.Extra.PaintSeed,
			Stickers:       stickers,
			OfferID:        *item.Extra.OfferId,
			ItemID:         item.ItemId,
			Timestamp:      time.Now(),
		}

		rawFloatPartValue := item.Extra.FloatPartValue
		if rawFloatPartValue != nil {
			floatPartValue, err := models.ParseFloatPartValue(string(*rawFloatPartValue))
			if err != nil {
				return nil, err
			}
			i.FloatPartValue = &floatPartValue
		}
		rawExterior := item.Extra.Exterior
		if rawExterior != nil {
			exterior, err := models.ParseExterior(string(*rawExterior))
			if err != nil {
				return nil, err
			}
			i.Exterior = &exterior
		}

		rawPhase := item.Extra.Phase
		if rawPhase != nil {
			phase, err := models.ParsePhase(string(*rawPhase))
			if err != nil {
				return nil, err
			}
			i.Phase = &phase
		}
		for _, filter := range filters {
			ok, err := filter(i)
			if err != nil {
				return nil, err
			}
			if !ok {
				continue Transform
			}
		}
		items = append(items, i)
	}
	return items, nil
}
