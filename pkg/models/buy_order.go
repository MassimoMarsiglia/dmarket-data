package models

type BuyOrder struct {
	MarketHashName string          `json:"market_hash_name"`
	FloatPartValue *FloatPartValue `json:"float_part_value"`
	Phase          *Phase          `json:"phase"`
	PaintSeed      *int            `json:"paint_seed"`
	Price          int             `json:"price"`
}
