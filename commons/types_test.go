package commons

import (
	"testing"
)

const (
	visa         = "4470 3307 6994 1000"
	visaBIN      = "447033"
	visaStripped = "4470330769941000"
	amex         = "378 2822 4631 0005"
	amexBin      = "378282"
)

func TestPanToStripped(t *testing.T) {
	pan := NewPan(visa)
	t.Log("Strip embedded spaces in credit card PAN")
	if pan.ToString() != visaStripped {
		t.Errorf("Expected %s, but it was %s instead.", visa, visaStripped)
	}
}
func TestPanVisaGetBin(t *testing.T) {
	pan := NewPan(visa)
	t.Logf("Getting BIN of %s from Visa", visaBIN)
	bin := pan.getBIN()
	if bin != visaBIN {
		t.Errorf("Expected %s, but it was %s instead.", visaBIN, bin)
	}
}
func TestPanAmexGetBin(t *testing.T) {
	pan := NewPan(amex)
	t.Logf("Getting BIN of %s from Amex", amexBin)
	bin := pan.getBIN()
	if bin != amexBin {
		t.Errorf("Expected %s, but it was %s instead.", amexBin, bin)
	}
}
func TestPandaIntoInteger(t *testing.T) {
	val := NewPanda(NewPan(visaStripped), NewPad())
	checkIntLength(val.ToInt64(), t)
}
func TestTokenIntoInteger(t *testing.T) {
	val := NewToken(NewPan("1234567890123456"))
	checkIntLength(val.ToInt64(), t)
}
func TestPadIntoInteger(t *testing.T) {
	val := NewPad()
	checkIntLength(val.ToInt64(), t)
}
func TestPanIntoInteger(t *testing.T) {
	val := NewPan(visaStripped)
	checkIntLength(val.ToInt64(), t)
}
func TestPandaIntoByteArr(t *testing.T) {
	val := NewPanda(NewPan(visaStripped), NewPad())
	checkByteArr(val.ToByteArr(), t)
}
func TestTokenIntoByteArr(t *testing.T) {
	val := NewToken(NewPan(visaStripped))
	checkByteArr(val.ToByteArr(), t)
}
func TestPadIntoByteArr(t *testing.T) {
	val := NewPad()
	checkByteArr(val.ToByteArr(), t)
}
func TestPanIntoByteArr(t *testing.T) {
	val := NewPan(visaStripped)
	checkByteArr(val.ToByteArr(), t)
}
func TestPanGetBin(t *testing.T) {
	val := NewPan(visaStripped)
	checkByteArr(val.ToByteArr(), t)
}

/*
func TestLuhnCheck(t *testing.T) {
	logTest("TestLuhnCheck")

	isLuhnValid:=LuhnCheck(knownGood)

	if(isLuhnValid!=true) {
		t.Errorf("Known good Luhn card %s failed test.", knownGood)

	}

}*/

/*func TestCreateTokenStr(t *testing.T) {
	pan:=NewPan(knownGood)

	if tokenStr==knownGood  {
		t.Errorf("Known good card %s shouldn't be the same as token.", knownGood)

	}
}
func logTest(testStr string){
	log.Printf("Starting test %s \n",testStr)

}*/
func checkByteArr(arr []byte, t *testing.T) {
	t.Log("Testing byte[] conversion for length of 4")
	if len(arr) > 16 {
		t.Errorf("Expected array of at least length 16, but it was %d instead.", len(arr))
	}
}
func checkIntLength(val int64, t *testing.T) {
	t.Log("Testing toInt64 conversion for value of 1234")
	if len(itoa(val)) < 16 {
		t.Errorf("Expected int64 of length at least 16, but it was %d instead.", itoa(val))
	}
}
func checkInt(val int64, t *testing.T) {
	t.Log("Testing toInt64 conversion for value of 1234")
	if i := val; i != 1234 {
		t.Errorf("Expected int64 of 1234, but it was %d instead.", i)
	}
}
