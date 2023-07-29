package pipelines

import . "github.com/enjekt/shannon-engine/models"

func NewPipeline() Pipeline {
	return &pipeline{}
}

type Pipeline interface {
	Add(pf PaletteFunc) Pipeline
	Execute(p Palette) Palette
}

type pipeline struct {
	head, tail chan Palette
}

func (p *pipeline) Execute(pa Palette) Palette {
	p.head <- pa
	return <-p.tail
}

func (p *pipeline) Add(pf PaletteFunc) Pipeline {

	if p.head == nil {
		p.head = make(chan Palette, 1)
		p.tail = p.head
	}
	outputChannel := make(chan Palette, 1)

	go pf(p.tail, outputChannel)
	p.tail = outputChannel

	return p
}
