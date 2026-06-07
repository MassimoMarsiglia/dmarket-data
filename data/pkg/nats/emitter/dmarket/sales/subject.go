package sales

import (
	"fmt"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

func (e *Emitter) Subject(s models.Sale) string {
	fpv := "nil"
	if s.FloatPartValue != nil {
		fpv = s.FloatPartValue.String()
	}

	phase := "nil"
	if s.Phase != nil {
		phase = s.Phase.String()
	}

	return fmt.Sprintf("dmarket.sales.%s.%s.%s", s.MarketHashName, fpv, phase)
}
