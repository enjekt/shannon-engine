package shannon_engine

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestPipeline(t *testing.T) {
	pan := Pan("5513746525703556")
	data := &palette{Pan: pan}
	encipherPipeline := &pipeline{}
	encipherPipeline.Add(CompactAndStripPanFunc).Add(CreatePadFunc).Add(EncipherFunc).Add(TokenFunc(6, 4))

	data = encipherPipeline.Execute(data)

	logStr := data.LogsToString()
	log.Println("Returned log string:\n", logStr)

	assert.NotEqual(t, data.Pan, data.PaddedPan)
	assert.NotEqual(t, data.Pan, data.Pad)
	assert.NotEqual(t, data.Pad, data.PaddedPan)

	decipherPipeline := &pipeline{}
	decipherPipeline.Add(DecipherFunc)

	data.Pan = ""
	log.Println(data.PaddedPan, data.Pad)
	decipherPipeline.Execute(data)

	assert.Equal(t, pan, data.Pan)

}
