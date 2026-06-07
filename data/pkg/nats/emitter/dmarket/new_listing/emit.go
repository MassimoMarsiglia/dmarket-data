package newlisting

import (
	"encoding/json"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

func (e *Emitter) Emit(i models.Item) error {
	subject := e.Subject(i)
	data, err := json.Marshal(i)
	if err != nil {
		return fmt.Errorf("marshal item: %w", err)
	}
	if err := e.nc.Publish(subject, data); err != nil {
		return fmt.Errorf("publish %s: %w", subject, err)
	}
	return nil
}
