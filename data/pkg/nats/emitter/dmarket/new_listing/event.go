package newlisting

import (
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

// NewNewListingEvent creates a NewListingEvent from a domain model Item.
// It maps the item fields into the event structure and attaches metadata.
func NewNewListingEvent(i models.Item) NewListingEvent {

	var owner *emitter.ItemOwner
	if i.Owner != nil {
		owner = &emitter.ItemOwner{
			OwnerId:  &i.Owner.ID,
			WalletId: &i.Owner.Wallet,
		}
	}

	var floatPartValue *emitter.FloatPartValue
	if i.FloatPartValue != nil {
		floatPartValue = utils.Ptr(emitter.FloatPartValue(i.FloatPartValue.String()))
	}

	var phase *emitter.Phase
	if i.Phase != nil {
		phase = utils.Ptr(emitter.Phase(i.Phase.String()))
	}

	numStickers := len(i.Stickers)
	stickers := make([]emitter.Sticker, 0, numStickers)
	if numStickers > 0 {
		for idx := range i.Stickers {
			stickers = append(stickers, emitter.Sticker{
				ImageUrl: i.Stickers[idx].Image,
				Name:     i.Stickers[idx].Name,
			})
		}
	}

	item := emitter.Item{
		ItemId:         i.ItemID,
		ImageUrl:       i.Image,
		OfferId:        i.OfferID,
		SkinId:         i.SkinID,
		Price:          int(i.Price),
		Owner:          owner,
		FloatValue:     i.Float,
		FloatPartValue: floatPartValue,
		MarketHashName: i.MarketHashName,
		PaintSeed:      i.PaintSeed,
		Market:         i.Market.String(),
		Phase:          phase,
		Timestamp:      i.Timestamp,
		Stickers:       stickers,
	}

	return NewListingEvent{
		Item:           item,
		Source:         "dmarket/new_listing",
		EventTimestamp: time.Now(),
	}
}
