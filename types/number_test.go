package types

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var initialStr = "1234567890123456"
var initialUint uint64 = 1234567890123456
var expectedSetStr = "98765432109876"
var newInt uint64 = 98765432109876

func TestPalette(t *testing.T) {
	p := NewPalette()
	p.GetPan().Set(initialStr)
	testVals(t, p.GetPan())
	p.GetPad().Set(initialStr)
	testVals(t, p.GetPad())
	p.GetPaddedPan().Set(initialStr)
	testVals(t, p.GetPaddedPan())
	p.GetToken().Set(initialStr)
	testVals(t, p.GetToken())

}
func TestPaletteMarshaling(t *testing.T) {
	p := NewPalette()
	p.GetPan().Set("1")
	p.GetPad().Set("2")
	p.GetPaddedPan().Set("3")
	p.GetToken().Set("4")

	// Marshal to json
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStr := string(j)
	// Print Json
	fmt.Printf("Json: %s", jsonStr)

}
func TestFieldTypes(t *testing.T) {
	testVals(t, NewPad(initialStr))
	testVals(t, NewPan(initialStr))
	testVals(t, NewToken(initialStr))
	testVals(t, NewPaddedPan(initialStr))
}

func testVals(t *testing.T, field Field) {
	assert.Equal(t, initialStr, field.String())
	assert.Equal(t, initialUint, field.ToUint64())
	field.SetUint64(newInt)
	assert.Equal(t, expectedSetStr, field.String())
	assert.Equal(t, newInt, field.ToUint64())

}
