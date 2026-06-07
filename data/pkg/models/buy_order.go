package models

import "time"

type BuyOrder struct {
	Market         MarketType      `json:"market"`
	MarketHashName string          `json:"market_hash_name"`
	SkinID         string          `json:"skin_id"`
	Depth          int             `json:"depth"`
	FloatPartValue *FloatPartValue `json:"float_part_value"`
	Phase          *Phase          `json:"phase"`
	PaintSeed      *int            `json:"paint_seed"`
	Price          int             `json:"price"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// return true if order is the same order
func (b *BuyOrder) IsEqual(inc BuyOrder) bool {
	if b.Phase == inc.Phase &&
		b.FloatPartValue == inc.FloatPartValue &&
		b.PaintSeed == inc.PaintSeed &&
		b.Price == inc.Price &&
		b.MarketHashName == inc.MarketHashName &&
		b.Depth == inc.Depth {
		return true
	}
	return false
}
