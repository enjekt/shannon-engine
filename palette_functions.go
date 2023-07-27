package shannon_engine

type PaletteFunc = func(input, output chan *palette)

var CompactAndStripPanFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Compact and strip..")
		evt.Pan = Pan(CompactAndStrip(string(evt.Pan)))
		output <- evt
	}

}

var CreatePadFunc = func(input, output chan *palette) {
	padChan := CreateNumberPump(16, 100)
	for evt := range input {
		evt.Log("Create Pad...")
		evt.Pad = Pad(<-padChan)
		output <- evt
	}

}

var EncipherFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Encipher the Pan...")
		evt.PaddedPan = Encipher(evt.Pan, evt.Pad)
		output <- evt
	}

}

var DecipherFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Decipher the Pan...")
		evt.Pan = Decipher(evt.PaddedPan, evt.Pad)
		output <- evt
	}

}

func TokenFunc(binLength, lastLength int) PaletteFunc {
	fillStrChan := CreateNumberPump(6, 100)
	pf := func(input, output chan *palette) {
		for evt := range input {
			evt.Log("Create Token...")

			pan := string(evt.Pan)
			bin := Bin(pan, binLength)
			last := Last(pan, lastLength)
			fill := <-fillStrChan
			evt.Token = Token(bin + fill + last)
			//Loop to get an non-PAN version
			for LuhnCheck(string(evt.Token)) {
				evt.Log("Illegal token created. Retry.")
				fill := <-fillStrChan
				evt.Token = Token(bin + fill + last)

			}
			output <- evt

		}
	}
	return pf
}
