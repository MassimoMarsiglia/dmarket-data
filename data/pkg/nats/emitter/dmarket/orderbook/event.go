package orderbook

import (
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

// NewOrderbookEvent creates an OrderbookEvent from a domain model BuyOrder.
func NewOrderbookEvent(order models.BuyOrder) OrderbookEvent {
	var floatPartValue *emitter.FloatPartValue
	if order.FloatPartValue != nil {
		floatPartValue = utils.Ptr(emitter.FloatPartValue(order.FloatPartValue.String()))
	}

	var phase *emitter.Phase
	if order.Phase != nil {
		phase = utils.Ptr(emitter.Phase(order.Phase.String()))
	}

	return OrderbookEvent{
		Item: emitter.BuyOrder{
			Market:         order.Market.String(),
			MarketHashName: order.MarketHashName,
			SkinId:         order.SkinID,
			Depth:          order.Depth,
			FloatPartValue: floatPartValue,
			Phase:          phase,
			PaintSeed:      order.PaintSeed,
			Price:          order.Price,
			UpdatedAt:      order.UpdatedAt,
		},
		Source:         "dmarket/orderbook",
		EventTimestamp: time.Now(),
	}
}
