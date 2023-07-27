package shannon_engine

import (
	"encoding/json"
	"strconv"
)

// TODO better type mechanism for use with methods.
type numberStr struct{ val string }
type NumberString interface {
	Set(string) NumberString
	String() string
	ToUint64() uint64
	SetUint64(uint64) NumberString
	MarshalJSON() ([]byte, error)
}
type Token interface {
	NumberString
}
type token struct {
	numberStr
}

func NewToken() Token {
	return &token{}
}

type PaddedPan interface {
	NumberString
}
type paddedPan struct {
	numberStr
}

func NewPaddedPan() PaddedPan {
	return &paddedPan{}
}

type Pad interface {
	NumberString
}
type pad struct {
	numberStr
}

func NewPad() Pad {
	return &pad{}
}

type Pan interface {
	NumberString
}
type pan struct {
	numberStr
}

func NewPan() Pan {
	return &pan{}
}
func (t *numberStr) Set(val string) NumberString {
	t.val = val
	return t
}
func (t *numberStr) SetUint64(val uint64) NumberString {
	t.val = strconv.FormatUint(val, 10)
	return t
}
func (t *numberStr) String() string {
	return t.val
}

func (t *numberStr) ToUint64() uint64 {
	return toUint64(t.val)
}

func (t *numberStr) MarshalJSON() ([]byte, error) {
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
