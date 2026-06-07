package models

type (
	Sale struct {
		Market           MarketType      `json:"market"`
		MarketHashName   string          `json:"market_hash_name"`
		SkinID           string          `json:"skin_id"`
		Price            int64           `json:"price"`
		FloatPartValue   *FloatPartValue `json:"float_part_value"`
		Phase            *Phase          `json:"phase"`
		Float            *float32        `json:"float_value"`
		PaintSeed        *int            `json:"paint_seed"`
		Stickers         []Sticker       `json:"stickers"`
		Exterior         *Exterior       `json:"exterior"`
		SaleType         SaleType        `json:"sale_type"`
		StickerSupported bool            `json:"stickers_supported"`
	}
)
