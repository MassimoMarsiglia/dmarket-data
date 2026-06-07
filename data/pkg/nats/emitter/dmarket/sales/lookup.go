package sales

import (
	"errors"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

var ErrInvalidMarketHashName = errors.New("invalid market hash name: ")

func (e *Emitter) Lookup(s *models.Sale) error {
	val, ok := e.lookup[s.MarketHashName]
	if !ok {
		return fmt.Errorf("%w %s", ErrInvalidMarketHashName, s.MarketHashName)
	}
	s.SkinID = val
	return nil
}
