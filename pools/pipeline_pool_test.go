package pools

import (
	"log"
	"testing"
)
import . "shannon-engine/pipelines"
import . "shannon-engine/types"

var testNumbers = []string{"5513 7465 2570 3556", "4532212776500464", "55 3221 2431868 196", "4716143551148112", "4716358667016165", "6011867865209833", "4916179771986533", "5515208833720309", "347850880061734"}

func TestPipelinePool(t *testing.T) {
	pool := constructBin6PlusLast4Pool()

	for _, numStr := range testNumbers {
		pipeline := pool.CheckOut()
		parsed := pipeline.Execute(NewPanPalette(numStr))
		log.Println(parsed)
		pool.CheckIn(pipeline)
	}
}

func constructBin6PlusLast4Pool() PipeLinePool {
	pool := NewPool(10)
	for i := 0; i < 9; i++ {
		pool.CheckIn(NewPipeline().Add(CompactAndStripPanFunc).Add(CreatePadFunc).Add(EncipherFunc).Add(TokenFunc(6, 4)))
	}

	return pool
}
