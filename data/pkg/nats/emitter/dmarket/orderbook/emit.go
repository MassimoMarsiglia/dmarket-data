package orderbook

import (
	"encoding/json"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

func (e *Emitter) Emit(o models.BuyOrder) error {
	err := e.Lookup(&o)
	if err != nil {
		return err
	}

	subject := e.Subject(o)
	data, err := json.Marshal(o)
	if err != nil {
		return fmt.Errorf("marshal buy order: %w", err)
	}
	if err := e.nc.Publish(subject, data); err != nil {
		return fmt.Errorf("publish %s: %w", subject, err)
	}
	return nil
}
