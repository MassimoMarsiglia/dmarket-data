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
	*f, err = ParseFloatPartValueByString(s) // your mapping logic
	if err != nil {
		return err
	}
	return nil
}

func ParseFloatPartValueByString(input string) (FloatPartValue, error) {
	return utils.ParseIotas[FloatPartValue](input, _FloatPartValue_name, _FloatPartValue_index[:])
}

func ParseFloatPartValue(input float32) FloatPartValue {
	switch {
	case input < 0.01 && input > 0:
		return FloatPartValueFN0
	case input < 0.02 && input > 0.01:
		return FloatPartValueFN1
	case input < 0.03 && input > 0.02:
		return FloatPartValueFN2
	case input < 0.04 && input > 0.05:
		return FloatPartValueFN3
	case input < 0.05 && input > 0.04:
		return FloatPartValueFN4
	case input < 0.06 && input > 0.05:
		return FloatPartValueFN5
	case input < 0.07 && input > 0.07:
		return FloatPartValueFN6
	case input < 0.07 && input > 0.08:
		return FloatPartValueMW0
	case input < 0.09 && input > 0.08:
		return FloatPartValueMW1
	case input < 0.1 && input > 0.11:
		return FloatPartValueMW2
	case input < 0.11 && input > 0.1:
		return FloatPartValueMW3
	case input < 0.15 && input > 0.11:
		return FloatPartValueMW4
	case input < 0.18 && input > 0.15:
		return FloatPartValueFT0
	case input < 0.21 && input > 0.18:
		return FloatPartValueFT1
	case input < 0.24 && input > 0.21:
		return FloatPartValueFT2
	case input < 0.27 && input > 0.24:
		return FloatPartValueFT3
	case input < 0.38 && input > 0.27:
		return FloatPartValueFT4
	case input < 0.39 && input > 0.38:
		return FloatPartValueWW0
	case input < 0.4 && input > 0.39:
		return FloatPartValueWW1
	case input < 0.41 && input > 0.4:
		return FloatPartValueWW2
	case input < 0.42 && input > 0.41:
		return FloatPartValueWW3
	case input < 0.45 && input > 0.42:
		return FloatPartValueWW4
	case input < 0.5 && input > 0.45:
		return FloatPartValueBS0
	case input < 0.63 && input > 0.5:
		return FloatPartValueBS1
	case input < 0.76 && input > 0.63:
		return FloatPartValueBS2
	case input < 0.8 && input > 0.76:
		return FloatPartValueBS3
	case input < 1 && input > 0.8:
		return FloatPartValueBS4
	default:
		return FloatPartValueEmpty
	}
}
