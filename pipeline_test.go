package shannon_engine

import (
	"log"
	"testing"
)

func TestPipeline(t *testing.T) {
	palette := &palette{Pan: "5513746525703556"}
	pipeline := &pipeline{}
	pipeline.Add(CompactAndStripPanFunc).Add(CreatePadFunc).Add(CreateTokenFunc).Add(PadPanFunc)
	pipeline.Execute(palette)
	logStr := palette.LogsToString()
	log.Println(logStr)
}
