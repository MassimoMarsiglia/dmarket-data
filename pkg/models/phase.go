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

func ParsePhase(input string) (Phase, error) {
	return utils.ParseIotas[Phase](input, _Phase_name, _Phase_index[:])
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
	*f, err = ParsePhase(s) // your mapping logic
	if err != nil {
		return err
	}
	return nil
}
