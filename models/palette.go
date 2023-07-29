package models

import "encoding/json"

func NewPalette() Palette {
	return &palette{Token: NewToken(""), PaddedPan: NewPaddedPan(""), Pad: NewPad(""), Pan: NewPan("")}
}
func NewPanPalette(pan string) Palette {
	return &palette{Token: NewToken(""), PaddedPan: NewPaddedPan(""), Pad: NewPad(""), Pan: NewPan(pan)}
}

type Palette interface {
	GetToken() Token
	GetPaddedPan() PaddedPan
	GetPad() Pad
	GetPan() Pan
	ToJSON() string
	Loggable
}

type palette struct {
	Token     `json:"Token"`
	PaddedPan `json:"PaddedPan"`
	Pad       `json:"Pad"`
	Pan       `json:"Pan"`
	LogData   `json:"-"`
}

func (p *palette) GetToken() Token {
	return p.Token
}

func (p *palette) GetPaddedPan() PaddedPan {
	return p.PaddedPan
}

func (p *palette) GetPad() Pad {
	return p.Pad
}

func (p *palette) GetPan() Pan {
	return p.Pan
}
func (p *palette) ToJSON() string {
	val, _ := json.Marshal(*p)
	return string(val)
}
