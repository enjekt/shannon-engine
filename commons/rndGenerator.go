package commons

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

// Buffered channel of random ints coming from a go subroutine closure
// Note almost all int handling in GoLang is uint64
var rndNumChannel = make(chan uint64, 1600)
var padPumpChannel = make(chan string, 100)
var sixDigitPumpChannel = make(chan string, 100)
var fiveDigitPumpChannel = make(chan string, 100)

func init() {
	initRndNumPump()
	initPadPump()
	initSixDigitPump()
	initFourDigitPump()
}

// Async random number pump prefills random numbers in separate process.
func initRndNumPump() {
	go func() {
		for {
			//0...9
			nBig, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				panic(err)
			}
			rndNumChannel <- nBig.Uint64()
		}
	}()
}

// Async random 16 digit pump pre-generates pads
func initPadPump() {
	go func() {
		for {
			padPumpChannel <- CreateRndDigitStr(16)
		}
	}()
}

func initSixDigitPump() {
	go func() {
		for {
			sixDigitPumpChannel <- CreateRndDigitStr(6)
		}
	}()
}

func initFourDigitPump() {
	go func() {
		for {
			fiveDigitPumpChannel <- CreateRndDigitStr(5)
		}
	}()
}

func CreateRandomSixteenDigitStr() string {
	return <-padPumpChannel
}

// Used for creating middle of Visa/MC tokens
func CreateRndSixDigitStr() string {
	return <-sixDigitPumpChannel
}
func CreateRndFiveDigitStr() string {
	return <-fiveDigitPumpChannel
}
func CreateRndDigitStr(len int) string {
	return string(getRndByteArr(len))
}

func fetchRndNum() uint64 {
	return <-rndNumChannel
}

func getRndByteArr(len int) []byte {

	var rnd []byte
	for i := 0; i < len; i++ {
		rnd = strconv.AppendUint(rnd, fetchRndNum(), 10)
	}
	return rnd
}
