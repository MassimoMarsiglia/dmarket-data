package orderbook

import (
	"errors"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

var ErrInvalidMarketHashName = errors.New("invalid market hash name: ")

func (e *Emitter) Lookup(o *models.BuyOrder) error {
	val, ok := e.lookup[o.MarketHashName]
	if !ok {
		return fmt.Errorf("%w %s", ErrInvalidMarketHashName, o.MarketHashName)
	}
	o.SkinID = val
	return nil
}
