package shannon_engine

func New() Pipeline {
	return &pipeline{}
}

type Pipeline interface {
	Add(pf PaletteFunc) Pipeline
	Execute(p *palette) *palette
}

type pipeline struct {
	head, tail chan *palette
}

func (p *pipeline) Execute(pa *palette) *palette {
	p.head <- pa
	return <-p.tail
}

func (p *pipeline) Add(pf PaletteFunc) Pipeline {

	if p.head == nil {
		p.head = make(chan *palette, 1)
		p.tail = p.head
	}
	outputChannel := make(chan *palette, 1)

	go pf(p.tail, outputChannel)
	p.tail = outputChannel

	return p
}
