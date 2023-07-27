package types

import (
	"github.com/stretchr/testify/assert"
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
