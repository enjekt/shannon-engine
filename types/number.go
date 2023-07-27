package types

import (
	"encoding/json"
	"strconv"
)

type Field interface {
	MarshalJSON() ([]byte, error)
	Set(val string)
	SetUint64(val uint64)
	ToUint64() uint64
	String() string
}
type fieldType struct {
	string
}

// //
type Pan interface {
	Field
}

type pan struct {
	fieldType
}

func NewPan(padStr string) Pad {
	return &pan{fieldType{padStr}}
}

type Pad interface {
	Field
}

type pad struct {
	fieldType
}

func NewPad(val string) Pad {
	return &pad{fieldType{val}}
}

type PaddedPan interface {
	Field
}

type paddedPan struct {
	fieldType
}

func NewPaddedPan(val string) Pad {
	return &paddedPan{fieldType{val}}
}

type Token interface {
	Field
}

type token struct {
	fieldType
}

func NewToken(val string) Pad {
	return &token{fieldType{val}}
}

func (t *fieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *fieldType) Set(val string) {
	t.string = val
}
func (t *fieldType) SetUint64(val uint64) {
	t.string = strconv.FormatUint(val, 10)
}
func (t *fieldType) String() string {
	return t.string
}

func (t *fieldType) ToUint64() uint64 {
	i, _ := strconv.ParseUint(t.string, 10, 64)
	return i
}
