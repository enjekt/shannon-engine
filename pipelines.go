package shannon_engine

type PadPipeline interface {
	Execute(palette Palette)
}

type padPipeline struct {
}

func (p padPipeline) Execute(palette Palette) {

}
