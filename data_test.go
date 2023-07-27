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
	palette.Token.Set("eenie")
	palette.PaddedPan.Set("meenie")
	palette.Pan.Set("minnie")
	palette.Pad.Set("moe")

	assert.Equal(t, "eenie", palette.Token.String())

	// Marshal to json
	j, err := json.Marshal(palette)
	if err != nil {
		log.Fatal(err)
	}
	jsonStr := string(j)
	// Print Json
	fmt.Printf("Json: %s", jsonStr)
	assert.True(t, strings.Contains(jsonStr, "\"token\":\"eenie\""))
}
