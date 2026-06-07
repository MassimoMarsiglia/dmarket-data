package newlisting

import (
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

func (e *Emitter) Subject(i models.Item) string {
	fpv := "nil"
	if i.FloatPartValue != nil {
		fpv = i.FloatPartValue.String()
	}

	phase := "nil"
	if i.Phase != nil {
		phase = i.Phase.String()
	}

	return fmt.Sprintf("dmarket.new_listing.%s.%s.%s", i.MarketHashName, fpv, phase)
}
