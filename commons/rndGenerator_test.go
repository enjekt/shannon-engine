package commons

import (
	"fmt"
	"log"
	"testing"
)

func TestRndStr(t *testing.T) {
	logTest("TestRndStr")
	for i := 0; i < 1000; i++ {
		rndStr := CreateRndDigitStr(10)

		//log.Printf("String generator line %d created %s of length \n", i, rndStr)
		if len(rndStr) != 10 {
			t.Errorf("Expected len %d string but received %s", 10, rndStr)
		}
	}
}
func TestGetPad(t *testing.T) {
	logTest("TestGetPad")
	for i := 0; i < 1000; i++ {
		rndPad := NewPad()
		//log.Printf("String generator line %d created %s of length \n", i, rndStr)
		//fmt.Println(rndStr)
		if rndPad.GetLength() != 16 {
			t.Errorf("Expected len %d string but received %s", 16, rndPad.ToString())
		}
	}
}
func TestGetRndSixDigitStr(t *testing.T) {
	logTest("TestGetRndFiveDigitStr")
	for i := 0; i < 10000; i++ {
		rndStr := CreateRndSixDigitStr()
		fmt.Println(rndStr)
		//log.Printf("String generator line %d created %s of length \n", i, rndStr)
		if len(rndStr) != 6 {
			t.Errorf("Expected len %d string but received len %d", 6, len(rndStr))
		}
	}
}
func TestGetRndFiveDigitStr(t *testing.T) {
	logTest("TestGetRndFiveDigitStr")
	for i := 0; i < 1000; i++ {
		rndStr := CreateRndFiveDigitStr()
		//log.Printf("String generator line %d created %s of length \n", i, rndStr)
		if len(rndStr) != 5 {
			t.Errorf("Expected len %d string but received %s", 5, rndStr)
		}
	}
}

func logTest(testStr string) {
	log.Printf("Starting test %s \n", testStr)

}
