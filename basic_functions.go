package shannon_engine

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

// Async random number pump prefills random numbers in separate process.

var CreateNumberPump = func(numberOfDigits, channelSize int) chan string {
	rndNumChannel := initRndNumPump(channelSize * numberOfDigits)
	numberStrChannel := make(chan string, channelSize)
	go func() {
		for {
			var rnd []byte
			for i := 0; i < numberOfDigits; i++ {
				rnd = strconv.AppendUint(rnd, <-rndNumChannel, 10)
			}
			numberStrChannel <- string(rnd)
		}
	}()
	return numberStrChannel
}

func initRndNumPump(channelSize int) chan uint64 {
	var rndNumChannel = make(chan uint64, channelSize)

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
	return rndNumChannel
}

var Encipher = func(pan Pan, pad Pad) string {
	xor := toUint64(string(pan)) ^ toUint64(string(pad))
	return strconv.FormatUint(xor, 10)
}

var Decipher = func(paddedPan, pad string) string {
	xor := toUint64(paddedPan) ^ toUint64(pad)
	return strconv.FormatUint(xor, 10)

}

func toUint64(bstr string) uint64 {
	var num uint64
	if i, err := strconv.ParseUint(bstr, 10, 64); err != nil {
		panic(err)
	} else {
		num = i
	}
	return num
}

type basicFunction = func(string) string

var CompactAndStrip = func(inputStr string) string {
	return strings.ReplaceAll(inputStr, " ", "")
}

var CreatePad = func(pan string) string {
	return "111111111111"
}

var Bin6 = func(numberStr string) string {
	return bankId(numberStr, 6)
}

var Bin8 = func(numberStr string) string {
	return bankId(numberStr, 8)
}

var Bin2 = func(numberStr string) string {
	return bankId(numberStr, 6)
}

func bankId(numberStr string, idx int) string {
	return numberStr[0:idx]
}

var LastTwo = func(numberStr string) string {
	return numberStr[len(numberStr)-2:]
}

var LuhnCheck = func(toCheck string) bool {
	toCheck = strings.ReplaceAll(toCheck, " ", "")
	checkInt, err := strToint64(CompactAndStrip(toCheck))
	if err != nil {
		return false
	}
	return LuhnCheckInt64(checkInt)
}

var LuhnCheckInt64 = func(toCheck int64) bool {
	return (toCheck%10+checkLuhnSum(toCheck/10))%10 == 0
}

var LastTwoInt64 = func(number int64) int64 {
	lastFour, _ := strToint64(LastTwo(int64ToStr(number)))
	return lastFour
}

// Helpers
func checkLuhnSum(number int64) int64 {
	var luhn int64

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}

func strToint64(toConvert string) (int64, error) {
	return strconv.ParseInt(toConvert, 10, 64)
}

func int64ToStr(toConvert int64) string {
	return strconv.FormatInt(int64(toConvert), 10)
}
