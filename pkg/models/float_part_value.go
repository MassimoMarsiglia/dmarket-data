package models

import (
	"encoding/json"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

type FloatPartValue int

//go:generate stringer -type=FloatPartValue -output=floatpartvalue_string.go -linecomment
const (
	FloatPartValueEmpty FloatPartValue = iota //Unknown
	FloatPartValueBS0                         //BS-0
	FloatPartValueBS1                         //BS-1
	FloatPartValueBS2                         //BS-2
	FloatPartValueBS3                         //BS-3
	FloatPartValueBS4                         //BS-4
	FloatPartValueFN0                         //FN-0
	FloatPartValueFN1                         //FN-1
	FloatPartValueFN2                         //FN-2
	FloatPartValueFN3                         //FN-3
	FloatPartValueFN4                         //FN-4
	FloatPartValueFN5                         //FN-5
	FloatPartValueFN6                         //FN-6
	FloatPartValueFT0                         //FT-0
	FloatPartValueFT1                         //FT-1
	FloatPartValueFT2                         //FT-2
	FloatPartValueFT3                         //FT-3
	FloatPartValueFT4                         //FT-4
	FloatPartValueMW0                         //MW-0
	FloatPartValueMW1                         //MW-1
	FloatPartValueMW2                         //MW-2
	FloatPartValueMW3                         //MW-3
	FloatPartValueMW4                         //MW-4
	FloatPartValueWW0                         //WW-0
	FloatPartValueWW1                         //WW-1
	FloatPartValueWW2                         //WW-2
	FloatPartValueWW3                         //WW-3
	FloatPartValueWW4                         //WW-4
)

func (f FloatPartValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *FloatPartValue) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var err error
	*f, err = ParseFloatPartValue(s) // your mapping logic
	if err != nil {
		return err
	}
	return nil
}

func ParseFloatPartValue(input string) (FloatPartValue, error) {
	return utils.ParseIotas[FloatPartValue](input, _FloatPartValue_name, _FloatPartValue_index[:])
}
