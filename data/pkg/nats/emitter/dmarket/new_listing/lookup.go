package newlisting

import (
	"errors"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

var ErrInvalidMarketHashName = errors.New("invalid market hash name: ")

func (e *Emitter) Lookup(i *models.Item) error {
	val, ok := e.lookup[i.MarketHashName]
	if !ok {
		return fmt.Errorf("%w %s", ErrInvalidMarketHashName, i.MarketHashName)
	}
	i.SkinID = val
	return nil
}
