package pipelines

import (
	"crypto/rand"
	"math/big"
	. "shannon-engine/types"
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

var Encipher = func(p Palette) {
	p.GetPaddedPan().SetUint64(p.GetPan().ToUint64() ^ p.GetPad().ToUint64())
}

var Decipher = func(p Palette) {
	p.GetPan().SetUint64(p.GetPaddedPan().ToUint64() ^ p.GetPad().ToUint64())
}

// Deprecate
//func ToUint64(bstr string) uint64 {
//	var num uint64
//	if i, err := strconv.ParseUint(bstr, 10, 64); err != nil {
//		panic(err)
//	} else {
//		num = i
//	}
//	return num
//}

type BasicFunction = func(string) string

var CompactAndStrip = func(inputStr string) string {
	return strings.ReplaceAll(inputStr, " ", "")
}

func Bin(numberStr string, length int) string {
	return numberStr[0:length]
}

var Last = func(numberStr string, length int) string {
	return numberStr[len(numberStr)-length:]
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
