package models

import (
	"encoding/json"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils"
)

type Phase int

//go:generate stringer -type=Phase -output=phase_string.go -linecomment

const (
	Doppler_Unknown    Phase = iota // Unknown
	Doppler_None                    //
	Doppler_BlackPearl              // black-pearl
	Doppler_Emerald                 // emerald
	Doppler_Phase1                  // phase-1
	Doppler_Phase2                  // phase-2
	Doppler_Phase3                  // phase-3
	Doppler_Phase4                  // phase-4
	Doppler_Ruby                    // ruby
	Doppler_Sapphire                // sapphire
)

func ParsePhase(market MarketType, input string) (Phase, error) {
	switch market {
	case UNKNOWN:
		return utils.ParseIotas[Phase](input, _Phase_name, _Phase_index[:])
	case DMARKET:
		return utils.ParseIotas[Phase](input, _Phase_name, _Phase_index[:])
	case BUFF:
		return parsePhaseBuff(input), nil
	}
	return utils.ParseIotas[Phase](input, _Phase_name, _Phase_index[:])
}

func parsePhaseBuff(input string) Phase {
	switch input {
	case "P1":
		return Doppler_Phase1
	case "P2":
		return Doppler_Phase2
	case "P3":
		return Doppler_Phase3
	case "P4":
		return Doppler_Phase4
	case "Sapphire":
		return Doppler_Sapphire
	case "Ruby":
		return Doppler_Ruby
	case "Emerald":
		return Doppler_Emerald
	case "Black_Pearl":
		return Doppler_BlackPearl
	default:
		return Doppler_Unknown
	}
}

func (p Phase) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (f *Phase) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var err error
	*f, err = ParsePhase(UNKNOWN, s) // your mapping logic
	if err != nil {
		return err
	}
	return nil
}
