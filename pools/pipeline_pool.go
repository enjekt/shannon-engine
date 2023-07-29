package pools

import . "github.com/enjekt/shannon-engine/pipelines"

func NewPool(numberOfPipelines int) PipeLinePool {
	return &pipeLinePool{pool: make(chan Pipeline, numberOfPipelines)}
}

type PipeLinePool interface {
	CheckIn(pool Pipeline)
	CheckOut() Pipeline
}
type pipeLinePool struct {
	pool chan Pipeline
}

func (p *pipeLinePool) CheckIn(pipeline Pipeline) {
	p.pool <- pipeline
}

func (p *pipeLinePool) CheckOut() Pipeline {
	return <-p.pool
}
