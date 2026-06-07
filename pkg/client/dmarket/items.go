package dmarket

import (
	"strconv"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

type FilterFunc[T any] func(T) (bool, error)

func applyFilters[T any](item T, filters []FilterFunc[T]) (bool, error) {
	for i := range filters {
		filter := filters[i]
		ok, err := filter(item)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

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
			Market:         models.DMARKET,
			MarketHashName: item.Title,
			Image:          item.Image,
			Price:          price,
			Owner:          &owner,
			Float:          item.Extra.FloatValue,
			PaintSeed:      item.Extra.PaintSeed,
			Stickers:       stickers,
			OfferID:        *item.Extra.OfferId,
			ItemID:         item.ItemId,
			Timestamp:      time.Now(),
		}

		rawFloatPartValue := item.Extra.FloatPartValue
		if rawFloatPartValue != nil {
			floatPartValue, err := models.ParseFloatPartValueByString(string(*rawFloatPartValue))
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
			phase, err := models.ParsePhase(i.Market, string(*rawPhase))
			if err != nil {
				return nil, err
			}
			i.Phase = &phase
		}

		ok, err := applyFilters[models.Item](i, filters)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue Transform
		}

		items = append(items, i)
	}
	return items, nil
}
