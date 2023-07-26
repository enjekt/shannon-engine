package shannon_engine

import "encoding/json"

type MarshalTypes string
type Token MarshalTypes
type PaddedPan MarshalTypes
type Pad MarshalTypes
type Pan MarshalTypes

func (types *MarshalTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(types)
}

type Palette struct {
	Token
	PaddedPan
	Pad
	Pan
}
