package shannon_engine

import "encoding/json"

// TODO better type mechanism for use with methods.
type NumericString string
type Token NumericString
type PaddedPan NumericString
type Pad NumericString
type Pan NumericString

func (t NumericString) String() string {
	return string(t)
}

func (t NumericString) ToUint64() uint64 {
	return toUint64(string(t))
}

func (t NumericString) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
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
