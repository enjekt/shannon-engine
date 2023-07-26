package shannon_engine

import "encoding/json"

type FieldType string
type Token FieldType
type PaddedPan FieldType
type Pad FieldType
type Pan FieldType

func (t *FieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

func NewPalette() Palette {
	return &palette{}
}

type Palette interface {
	Loggable
}

type palette struct {
	LogData
	Token
	PaddedPan
	Pad
	Pan
}
