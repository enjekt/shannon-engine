package commons

import (
	"strconv"
	"strings"
	"log"
	"fmt"
)

type PV struct {
	ID string `json:"Id,omitempty"`
}
type Token struct{PV}
type Panda struct{PV}
type Pad struct {PV}
type Pan struct {PV}

type Result struct {
	Response string `json:"Id"`
}

type TokenPadMessage struct {
	Token string `json:"Token,omitempty"`
	Pad string `json:"Pad,omitempty"`
}
/*type TokenPadMsg struct {
	Token string `json:"ID,omitempty"`
	Pad string `json:"ID,omitempty"`
}*/

func NewToken(pan *Pan) *Token {

	bin := pan.GetBin()
	last4 := pan.GetLastFour()
	var tokenRand string
	var tokenStr string

	//We are not going to allow Luhn compliant tokens as that is a basic
	//distingushing characteristic of a cred card number
	for LuhnCheck(tokenStr) {
		if pan.GetLength() == lenOfVisaMC {
			tokenRand = CreateRndSixDigitStr()
		} else if pan.GetLength() == lenOfAmex {
			tokenRand = CreateRndFiveDigitStr()
		} else {
			log.Fatalf("Non-standard card length %d for token %s", pan.GetLength(), pan)
			break
		}

		tokenStr =bin + tokenRand + last4
	}
	token:=Token{}
	token.ID =tokenStr
	return &token
}


func NewPanda(pan *Pan, pad *Pad) *Panda {
	pandaInt:=pan.ToInt64() + pad.ToInt64()
	panda:=&Panda{}
	panda.ID = itoa(pandaInt)
	return panda
}

func NewPad() *Pad{
	p:=&Pad{}
	p.ID = CreateRandomSixteenDigitStr()
	return p
}
func NewPan(str string) *Pan{
	p:=&Pan{}
	p.ID =stripSpaces(str)
	return p
}

func InitPad(pandaStr string) *Pad {
	pad:=&Pad{}
	pad.ID = stripSpaces(pandaStr)
	return pad
}


func InitPanda(pandaStr string) *Panda {
	panda:=&Panda{}
	panda.ID = pandaStr
	return panda
}

func InitToken(tokenStr string) *Token {
	token:=Token{}
	token.ID =tokenStr
	return &token
}
func InitPan(panda *Panda, pad *Pad) *Pan {
	pan:=Pan{}
	pan.ID =itoa(panda.ToInt64()-pad.ToInt64())
	return &pan
}
func (self *Pan) GetBin() string {
	return self.ID[0:6]
}
func (self *TokenPadMessage)  ToString() string {
	return self.Token +"," + self.Pad
}
func (self *Pan)  GetLastFour() string {
	return self.ID[len(self.ID)-4:]
}
func (self *Pan)  GetPan() string {
	return self.ID
}
func (self *PV)  GetLength() int {
	return len(self.ID)
}

func stripSpaces(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func (self *PV) ToByteArr() []byte {
	return []byte(self.ID)
}
func (self *PV) ToString() string {
	return stripSpaces(self.ID)
}
func (self *PV) ToInt64() int64 {
	return toInt64(self.ID)
}
//BIN is the bank identifier at beginning of card of 6 digits
func (self *Pan) getBIN() string {
	return self.ID[0:6]
}

const (
	lenOfVisaMC = 16
	lenOfAmex   = 15
)

func atoi(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
func itoa(num int64) string {
	str := strconv.FormatInt(num, 10)
	return str
}


func LuhnCheck(tokenStr string) bool {
	var sum int64
	var value int64
	idx := len(tokenStr) // Start from the end of string
	alt := false

	for idx > 0 {
		idx--
		// Get value. Throws error if it isn't a digit
		value = atoi(tokenStr[idx : idx+1])
		if alt {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}
		sum += value
		alt = !alt // Toggle alt-flag
	}
	luhnCheck:=(sum % 10) == 0
	fmt.Println("Luhn check: ", luhnCheck)
	return luhnCheck

}
func toInt64(pv string) int64 {
	i, _ := strconv.ParseInt(string(pv), 10, 64)
	return i
}
