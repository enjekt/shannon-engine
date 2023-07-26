package shannon_engine

type paletteFunc = func(p *palette) *palette

var CompactAndStripPanFunc = func(p *palette) *palette {
	p.Log("Compact and Strip Pan ")
	return p
}

var CreatePadFunc = func(p *palette) *palette {
	p.Log("Create Pad ")
	return p
}

var PadPanFunc = func(p *palette) *palette {
	p.Log("Pad the Pan")
	return p
}

var CreateTokenFunc = func(p *palette) *palette {
	p.Log("Create Token")
	return p
}
