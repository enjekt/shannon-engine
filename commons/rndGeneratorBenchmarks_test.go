package commons

import (
	"testing"
)

func BenchmarkHundredThousandPads(b *testing.B) {

	for i:=0 ; i< b.N; i++ {
		NewPad()
	}
}
