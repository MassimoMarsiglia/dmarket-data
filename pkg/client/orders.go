package client

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

var ErrMissingOrders = errors.New("orders missing")

func OrderDepthFilter() FilterFunc[models.BuyOrder] {
	mu := sync.Mutex{}
	// market_hash_name, start_depth, order
	seenOrders := make(map[string]map[int]models.BuyOrder)

	return func(i models.BuyOrder) (bool, error) {
		mu.Lock()
		defer mu.Unlock()

		orderBook, ok := seenOrders[i.MarketHashName]
		if !ok {
			orderBook = make(map[int]models.BuyOrder, 0)
			orderBook[i.Depth] = i
			seenOrders[i.MarketHashName] = orderBook
			return true, nil
		}

		order, ok := orderBook[i.Depth]
		if !ok {
			orderBook[i.Depth] = i
			return true, nil
		}

		if i.IsEqual(order) {
			order.UpdatedAt = time.Now()
			return false, nil
		}

		orderBook[i.Depth] = i
		return true, nil
	}
}

func (e EntityGetOrderBookResponse) GetOrders(marketHashName string, filters ...FilterFunc[models.BuyOrder]) ([]models.BuyOrder, error) {
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
						err = nil
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
				ps := *attr.PaintSeed
				if ps != "any" && ps != "all" && ps != "none" {
					psi64, err := strconv.ParseInt(*attr.PaintSeed, 10, 16)
					if err != nil {
						return nil, err
					}
					paintSeed = utils.Ptr(int(psi64))
				}
			}

			d64, err := strconv.ParseInt(order.Liquidity, 10, 64)
			if err != nil {
				return nil, err
			}
			depth := int(d64)

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
				Depth:          depth,
				UpdatedAt:      time.Now(),
			})
		}
	}
	return orders, nil
}
