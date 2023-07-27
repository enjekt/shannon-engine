package shannon_engine

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestMarshalTypes(t *testing.T) {
	palette := &palette{}
	palette.Token = "eenie"
	palette.PaddedPan = "meenie"
	palette.Pan = "minnie"
	palette.Pad = "moe"

	assert.Equal(t, "eenie", palette.Token)

	// Marshal to json
	j, err := json.Marshal(palette)
	if err != nil {
		log.Fatal(err)
	}
	jsonStr := string(j)
	// Print Json
	fmt.Printf("Json: %s", jsonStr)
	assert.True(t, strings.Contains(jsonStr, "\"Token\":\"eenie\""))
}
