package listing

import (
	"encoding/json"
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

func (e *Emitter) Emit(i models.Item) error {
	err := e.Lookup(&i)
	if err != nil {
		return err
	}

	event := NewListingEvent(i)
	subject := e.Subject(i)
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}
	if err := e.nc.Publish(subject, data); err != nil {
		return fmt.Errorf("publish %s: %w", subject, err)
	}
	return nil
}
