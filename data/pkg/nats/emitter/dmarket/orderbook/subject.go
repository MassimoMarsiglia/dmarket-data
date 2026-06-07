package orderbook

import (
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

func (e *Emitter) Subject(o models.BuyOrder) string {
	fpv := "nil"
	if o.FloatPartValue != nil {
		fpv = o.FloatPartValue.String()
	}

	phase := "nil"
	if o.Phase != nil {
		phase = o.Phase.String()
	}

	return fmt.Sprintf("dmarket.orderbook.%s.%s.%s", o.MarketHashName, fpv, phase)
}
