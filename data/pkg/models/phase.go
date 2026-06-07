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

var phaseBuffMap = map[Phase]string{
	Doppler_Phase1:     "P1",
	Doppler_Phase2:     "P2",
	Doppler_Phase3:     "P3",
	Doppler_Phase4:     "P4",
	Doppler_Sapphire:   "Sapphire",
	Doppler_Ruby:       "Ruby",
	Doppler_Emerald:    "Emerald",
	Doppler_BlackPearl: "Black_Pearl",
}

var buffPhaseMap = map[string]Phase{
	"P1":          Doppler_Phase1,
	"P2":          Doppler_Phase2,
	"P3":          Doppler_Phase3,
	"P4":          Doppler_Phase4,
	"Sapphire":    Doppler_Sapphire,
	"Ruby":        Doppler_Ruby,
	"Emerald":     Doppler_Emerald,
	"Black_Pearl": Doppler_BlackPearl,
}

func parsePhaseBuff(input string) Phase {
	phase, ok := buffPhaseMap[input]
	if !ok {
		return Doppler_Unknown
	}
	return phase
}

func (p Phase) BuffString() string {
	phaseStr, ok := phaseBuffMap[p]
	if !ok {
		return ""
	}
	return phaseStr
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
