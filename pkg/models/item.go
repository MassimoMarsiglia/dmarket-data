package models

import "time"

type (
	Item struct {
		MarketType     MarketType      `json:"market"`
		MarketHashName string          `json:"market_hash_name"`
		Image          string          `json:"image_url"`
		Price          int64           `json:"price"`
		Owner          *Owner          `json:"owner"`
		FloatPartValue *FloatPartValue `json:"float_part_value"`
		Phase          *Phase          `json:"phase"`
		Float          *float32        `json:"float_value"`
		PaintSeed      *int            `json:"paint_seed"`
		Stickers       []Sticker       `json:"stickers"`
		Exterior       *Exterior       `json:"exterior"`
		OfferID        string          `json:"offer_id"`
		ItemID         string          `json:"item_id"`
		Timestamp      time.Time       `json:"timestamp"`
	}

	Sticker struct {
		Name  string `json:"sticker_market_hash_name"`
		Image string `json:"image_url"`
	}

	Owner struct {
		ID     string `json:"owner_id"`
		Wallet string `json:"wallet_id"`
	}
)
