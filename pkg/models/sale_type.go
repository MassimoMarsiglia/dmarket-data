package models

import (
	"encoding/json"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

type SaleType int

//go:generate stringer -type=SaleType -output=sale_type_string.go -linecomment

const (
	Unknown SaleType = iota // unknown
	Order                   // order
	Offer                   // sale
)

func (f SaleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *SaleType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var err error
	*f, err = ParseSaleType(s) // your mapping logic
	if err != nil {
		return err
	}
	return nil
}

func ParseSaleType(input string) (SaleType, error) {
	return utils.ParseIotas[SaleType](input, _FloatPartValue_name, _FloatPartValue_index[:])
}
