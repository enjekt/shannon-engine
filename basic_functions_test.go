package shannon_engine

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

var validNumberStr = []string{"5513746525703556", "4532212776500464", "5532212431868196", "4716143551148112", "4716358667016165", "6011867865209833", "4916179771986533", "5515208833720309", "347850880061734"}
var binSixes = []string{"551374", "453221", "553221", "471614", "471635", "601186", "491617", "551520", "347850"}
var lastTwo = []string{"56", "64", "96", "12", "65", "33", "33", "09", "34"}
var validNumbers = []int64{5513746525703556, 4532212776500464, 5532212431868196, 4716143551148112, 4716358667016165, 6011867865209833, 4916179771986533, 5515208833720309, 347850880061734}

func TestEncipherAndDecipher(t *testing.T) {
	otp := Pad("987654321087654")
	for _, numStr := range validNumberStr {
		pan := Pan(numStr)
		paddedPan := Encipher(pan, otp)
		//log.Println(pan, "!=", paddedPan)
		assert.NotEqual(t, pan, paddedPan)
		panAgain := Decipher(paddedPan, otp)
		assert.Equal(t, pan, panAgain)
	}
}

func TestRndStr(t *testing.T) {
	tenDigitRandom := CreateNumberPump(10, 20)
	for i := 0; i < 1000; i++ {
		rndStr := <-tenDigitRandom
		//log.Printf("String generator line %d created %s of length \n", i, rndStr)
		if len(rndStr) != 10 {
			t.Errorf("Expected len %d string but received %s", 10, rndStr)
		}
	}
}
func TestCompactAndStrip(t *testing.T) {
	spacedOut := "0000 1111 2222 3333 4444"
	fixed := CompactAndStrip(spacedOut)
	assert.Equal(t, fixed, "00001111222233334444")
}

func TestLuhnCheckInt64(t *testing.T) {

	for _, validNumber := range validNumbers {
		assert.True(t, LuhnCheckInt64(validNumber))
	}
}

func TestLuhnCheck(t *testing.T) {

	for _, validNumber := range validNumberStr {
		assert.True(t, LuhnCheck(validNumber))

	}
}

func TestLastTwo(t *testing.T) {
	for loc, number := range validNumberStr {
		val := Last(number, 2)
		assert.Equal(t, val, lastTwo[loc])

	}

}

func TestBin6(t *testing.T) {
	for loc, number := range validNumberStr {
		bin := Bin(number, 6)
		assert.Equal(t, bin, binSixes[loc])
	}
}

func createArrayPrint(vals []string) {
	log.Println("var someStr = []string{", strings.Join(vals, ","), "}")
}
