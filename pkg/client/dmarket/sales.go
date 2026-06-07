package dmarket

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
)

var ErrMissingSales = errors.New("sales missing")

func (t TradeGetLastSalesResponse) GetSales(title string, filters ...FilterFunc[models.Sale]) ([]models.Sale, error) {
	if t.Sales == nil {
		return nil, fmt.Errorf("t.sales is nil")
	}

	sales := make([]models.Sale, 0, len(*t.Sales))
	if t.Sales == nil {
		return nil, fmt.Errorf("%w in Last Sales response", ErrMissingOrders)
	}

	for _, s := range *t.Sales {

		pricef64, err := strconv.ParseFloat(*s.Price, 64)
		if err != nil {
			return nil, err
		}
		priceInt := int64(pricef64 * 100)

		sale := models.Sale{
			Market:         models.DMARKET,
			MarketHashName: title,
			Price:          priceInt,
		}

		attributes := *s.OfferAttributes
		if s.OfferAttributes == nil {
			var floatPartValue *models.FloatPartValue
			f, ok := attributes["floatValue"]
			if ok {
				f32, ok := f.(float32)
				if ok {
					sale.Float = &f32
					fpv := models.ParseFloatPartValue(f32)
					floatPartValue = &fpv
				}
			}
			sale.FloatPartValue = floatPartValue

			var paintSeed *int
			ps, ok := attributes["paintSeed"]
			if ok {
				pInt, ok := ps.(int)
				if ok {
					paintSeed = &pInt
				}
			}
			sale.PaintSeed = paintSeed
		}

		ok, err := applyFilters[models.Sale](sale, filters)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		sales = append(sales, sale)
	}
	return sales, nil
}
