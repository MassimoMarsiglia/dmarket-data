package models

import (
	"encoding/json"
	"strings"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

//go:generate stringer -type=Exterior -output=exterior_string.go -linecomment

type Exterior int

const (
	BattleScarred Exterior = iota //battle-scarred
	FactoryNew                    //factory new
	FieldTested                   //field-tested
	MinimalWear                   //minimal wear
	NotPainted                    //not painted
	WellWorn                      //well-worn
)

func ParseExterior(input string) (Exterior, error) {
	input = strings.TrimSpace(strings.ToLower(input))
	return utils.ParseIotas[Exterior](input, _Exterior_name, _Exterior_index[:])
}

func (e Exterior) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (f *Exterior) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var err error
	*f, err = ParseExterior(s) // your mapping logic
	if err != nil {
		return err
	}
	return nil
}
