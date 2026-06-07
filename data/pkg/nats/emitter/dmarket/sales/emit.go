package sales

import (
	"encoding/json"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

func (e *Emitter) Emit(s models.Sale) error {
	subject := e.Subject(s)
	data, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("marshal sale: %w", err)
	}
	if err := e.nc.Publish(subject, data); err != nil {
		return fmt.Errorf("publish %s: %w", subject, err)
	}
	return nil
}
