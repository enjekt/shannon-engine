package shannon_engine

type PaletteFunc = func(input, output chan *palette)

var CompactAndStripPanFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Compact and strip..")
		evt.Pan.Set(CompactAndStrip(evt.Pan.String()))
		output <- evt
	}

}

var CreatePadFunc = func(input, output chan *palette) {
	padChan := CreateNumberPump(16, 100)
	for evt := range input {
		evt.Log("Create pad...")
		evt.Pad = NewPad().Set(<-padChan)
		output <- evt
	}

}

var EncipherFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Encipher the pan...")
		evt.PaddedPan = Encipher(evt.Pan, evt.Pad)
		output <- evt
	}

}

var DecipherFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Decipher the pan...")
		evt.Pan = Decipher(evt.PaddedPan, evt.Pad)
		output <- evt
	}

}

func TokenFunc(binLength, lastLength int) PaletteFunc {
	fillStrChan := CreateNumberPump(6, 100)
	pf := func(input, output chan *palette) {
		for evt := range input {
			evt.Log("Create token...")

			pan := evt.Pan.String()
			bin := Bin(pan, binLength)
			last := Last(pan, lastLength)
			fill := <-fillStrChan
			evt.Token = NewToken().Set(bin + fill + last)
			//Loop to get an non-PAN version
			for LuhnCheck(evt.Token.String()) {
				evt.Log("Illegal token created. Retry.")
				fill := <-fillStrChan
				evt.Token = NewToken().Set(bin + fill + last)

			}
			output <- evt

		}
	}
	return pf
}
