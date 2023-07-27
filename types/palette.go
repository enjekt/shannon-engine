package types

func NewPalette() Palette {
	return &palette{Token: NewToken(""), PaddedPan: NewPaddedPan(""), Pad: NewPad(""), Pan: NewPan("")}
}

type Palette interface {
	GetToken() Token
	GetPaddedPan() PaddedPan
	GetPad() Pad
	GetPan() Pan
	Loggable
}

type palette struct {
	Token
	PaddedPan
	Pad
	Pan
	LogData
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
