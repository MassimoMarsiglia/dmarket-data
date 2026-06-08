package sales

import (
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

// NewSalesEvent creates a SalesEvent from a domain model Sale.
func NewSalesEvent(sale models.Sale) SalesEvent {
	var floatPartValue *emitter.FloatPartValue
	if sale.FloatPartValue != nil {
		floatPartValue = utils.Ptr(emitter.FloatPartValue(sale.FloatPartValue.String()))
	}

	var phase *emitter.Phase
	if sale.Phase != nil {
		phase = utils.Ptr(emitter.Phase(sale.Phase.String()))
	}

	var exterior *emitter.Exterior
	if sale.Exterior != nil {
		exterior = utils.Ptr(emitter.Exterior(sale.Exterior.String()))
	}

	var stickers *[]emitter.Sticker
	if sale.Stickers != nil {
		s := make([]emitter.Sticker, len(sale.Stickers))
		for i, v := range sale.Stickers {
			s[i] = emitter.Sticker{
				Name:     v.Name,
				ImageUrl: v.Image,
			}
		}
		stickers = &s
	}

	return SalesEvent{
		Item: emitter.Sale{
			Market:            sale.Market.String(),
			MarketHashName:    sale.MarketHashName,
			SkinId:            sale.SkinID,
			Price:             int(sale.Price),
			FloatPartValue:    floatPartValue,
			Phase:             phase,
			FloatValue:        sale.Float,
			PaintSeed:         sale.PaintSeed,
			Stickers:          stickers,
			Exterior:          exterior,
			SaleType:          emitter.SaleType(sale.SaleType.String()),
			StickersSupported: sale.StickerSupported,
		},
		Source:         "dmarket/sales",
		EventTimestamp: time.Now(),
	}
}
