package client

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

var ErrMissingOrders = errors.New("orders missing")

func (e EntityGetOrderBookResponse) GetOrders(marketHashName string, filters ...FilterFunc[any]) ([]models.BuyOrder, error) {
	orders := make([]models.BuyOrder, 0, len(e.Orders))

	if e.Orders == nil {
		return nil, fmt.Errorf("%w in order book response", ErrMissingOrders)
	}

	for _, order := range e.Orders {
		var attributes []EntityOrderBookAttribute
		if order.Attributes != nil && len(*order.Attributes) != 0 {
			attr := *order.Attributes
			attributes = attr
		}

		for _, attr := range attributes {
			var floatPartValue *models.FloatPartValue
			if attributes != nil && attr.FloatPartValue != nil {
				if *attr.FloatPartValue != "any" {
					fpv, err := models.ParseFloatPartValue(string(*attr.FloatPartValue))
					if err != nil {
						return nil, err
					}
					floatPartValue = &fpv
				}
			}

			var phase *models.Phase
			if attributes != nil && attr.PhaseTitle != nil {
				if *attr.PhaseTitle != "any" {
					p, err := models.ParsePhase(string(*attr.PhaseTitle))
					if err != nil {
						err = nil
					}
					phase = &p
				}
			}

			var paintSeed *int
			if attributes != nil && attr.PaintSeed != nil {
				if *attr.PaintSeed != "any" {
					psi64, err := strconv.ParseInt(*attr.PaintSeed, 10, 16)
					if err != nil {
						return nil, err
					}
					paintSeed = utils.Ptr(int(psi64))
				}
			}

			p64, err := strconv.ParseInt(order.Price, 10, 64)
			if err != nil {
				return nil, err
			}
			price := int(p64)

			orders = append(orders, models.BuyOrder{
				MarketHashName: marketHashName,
				FloatPartValue: floatPartValue,
				Phase:          phase,
				PaintSeed:      paintSeed,
				Price:          price,
			})
		}
	}
	return orders, nil
}
