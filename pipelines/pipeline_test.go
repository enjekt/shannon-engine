package pipelines

import (
	. "github.com/enjekt/shannon-engine/models"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestPipeline(t *testing.T) {
	validationPan := "5513746525703556"
	data := NewPalette()
	data.GetPan().Set("5513746525703556")
	encipherPipeline := &pipeline{}
	encipherPipeline.Add(CompactAndStripPanFunc).Add(CreatePadFunc).Add(EncipherFunc).Add(TokenFunc(6, 4))

	data = encipherPipeline.Execute(data)

	logStr := data.LogsToString()
	log.Println("Returned log string:\n", logStr)

	assert.NotEqual(t, data.GetPan().String(), data.GetPaddedPan().String())
	assert.NotEqual(t, data.GetPan().String(), data.GetPad().String())
	assert.NotEqual(t, data.GetPad().String(), data.GetPaddedPan().String())

	decipherPipeline := &pipeline{}
	decipherPipeline.Add(DecipherFunc)

	data.GetPan().Set("")
	//log.Println(data.PaddedPan, data.Pad)
	decipherPipeline.Execute(data)

	assert.Equal(t, validationPan, data.GetPan().String())

}
