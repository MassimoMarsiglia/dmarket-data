package listing

import (
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

// NewListingEvent creates a ListingEvent from a domain model Item.
func NewListingEvent(item models.Item) ListingEvent {
	var owner *emitter.ItemOwner
	if item.Owner != nil {
		owner = &emitter.ItemOwner{
			OwnerId:  &item.Owner.ID,
			WalletId: &item.Owner.Wallet,
		}
	}

	var floatPartValue *emitter.FloatPartValue
	if item.FloatPartValue != nil {
		floatPartValue = utils.Ptr(emitter.FloatPartValue(item.FloatPartValue.String()))
	}

	var phase *emitter.Phase
	if item.Phase != nil {
		phase = utils.Ptr(emitter.Phase(item.Phase.String()))
	}

	var exterior *emitter.Exterior
	if item.Exterior != nil {
		exterior = utils.Ptr(emitter.Exterior(item.Exterior.String()))
	}

	numStickers := len(item.Stickers)
	stickers := make([]emitter.Sticker, 0, numStickers)
	if numStickers > 0 {
		for idx := range item.Stickers {
			stickers = append(stickers, emitter.Sticker{
				ImageUrl: item.Stickers[idx].Image,
				Name:     item.Stickers[idx].Name,
			})
		}
	}

	return ListingEvent{
		Item: emitter.Item{
			ItemId:         item.ItemID,
			ImageUrl:       item.Image,
			OfferId:        item.OfferID,
			SkinId:         item.SkinID,
			Price:          int(item.Price),
			Owner:          owner,
			FloatValue:     item.Float,
			FloatPartValue: floatPartValue,
			MarketHashName: item.MarketHashName,
			PaintSeed:      item.PaintSeed,
			Market:         item.Market.String(),
			Phase:          phase,
			Timestamp:      item.Timestamp,
			Stickers:       stickers,
			Exterior:       exterior,
		},
		Source:         "buff/listing",
		EventTimestamp: time.Now(),
	}
}
