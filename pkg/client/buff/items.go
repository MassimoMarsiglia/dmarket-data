package buff

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	buffqueue "github.com/MassimoMarsiglia/dmarket-bot/pkg/models/buff_queue"
)

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

func (s SellOrderResponse) Items(item buffqueue.Item, marketHashName string, filters ...FilterFunc[models.Item]) ([]models.Item, error) {
	items := make([]models.Item, 0, len(s.Data.Items))

Transform:
	for i := range s.Data.Items {
		item := s.Data.Items[i]

		parts := strings.Split(item.Price, ".")
		partLen := len(parts)
		if partLen > 2 {
			return nil, fmt.Errorf("invalid number of parts: %s", len(parts))
		}

		var priceStr string
		if partLen == 2 {
			whole := parts[0]
			frac := parts[1]

			if len(frac) == 1 {
				frac += "0"
			} else if len(frac) > 2 {
				frac = frac[:2]
			}
			priceStr = whole + frac
		}

		if partLen == 1 {
			priceStr = parts[0]
		}

		price, err := strconv.Atoi(priceStr)
		if err != nil {
			return nil, err
		}
		_ = price

		stickers := make([]models.Sticker, 0, len(item.AssetInfo.Info.Stickers))
		for _, sticker := range item.AssetInfo.Info.Stickers {
			stickers = append(stickers, models.Sticker{
				Name:  sticker.Name,
				Image: sticker.Name,
			})
		}

		var float float32
		if item.AssetInfo.Paintwear != "" {
			f64, err := strconv.ParseFloat(item.AssetInfo.Paintwear, 32)
			if err != nil {
				return nil, err
			}
			float = float32(f64)
		}

		floatPartValue := models.ParseFloatPartValue(float)

		marketPlace := models.BUFF

		var phase *models.Phase
		phaseData := item.AssetInfo.Info.PhaseData
		if phaseData != nil {
			p, err := models.ParsePhase(marketPlace, string(phaseData.Name))
			if err != nil {
				return nil, err
			}
			phase = &p
		}

		i := models.Item{
			Market:         marketPlace,
			MarketHashName: marketHashName,
			Image:          item.AssetInfo.Info.IconUrl,
			Price:          int64(price),
			Float:          &float,
			Phase:          phase,
			PaintSeed:      &item.AssetInfo.Info.Paintseed,
			Stickers:       stickers,
			OfferID:        item.Id,
			ItemID:         item.AssetInfo.Assetid,
			Timestamp:      time.Now(),
			FloatPartValue: &floatPartValue,
		}

		for z := range filters {
			filter := filters[z]
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
