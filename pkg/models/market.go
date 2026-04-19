package models

import (
	"encoding/json"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

//go:generate stringer -type=MarketType -output=market_type_string.go -linecomment

const (
	UNKNOWN MarketType = iota // unknown
	BUFF                      //buff
	DMARKET                   //dmark
)

type MarketType int

func (f MarketType) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *MarketType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var err error
	*f, err = ParseMarketType(s)
	if err != nil {
		return err
	}
	return nil
}

func ParseMarketType(input string) (MarketType, error) {
	return utils.ParseIotas[MarketType](input, _MarketType_name, _MarketType_index[:])
}
