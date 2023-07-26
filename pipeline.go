package shannon_engine

type Pipeline interface {
	Execute(palette *palette)
	Add(pf paletteFunc) Pipeline
}

type pipeline struct {
	paletteFunctions []paletteFunc
}

// TODO Initial step to simply chain. Then we'll add the bidirectional
// channels and remove return type on function call.
func (p *pipeline) Execute(palette *palette) {
	for _, pf := range p.paletteFunctions {
		palette = pf(palette)
	}
}

func (p *pipeline) Add(pf paletteFunc) Pipeline {
	p.paletteFunctions = append(p.paletteFunctions, pf)
	return p
}
