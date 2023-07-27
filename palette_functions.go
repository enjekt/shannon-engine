package shannon_engine

import . "shannon-engine/types"

type PaletteFunc = func(input, output chan Palette)

var CompactAndStripPanFunc = func(input, output chan Palette) {
	for evt := range input {
		evt.Log("Compact and strip..")
		evt.GetPan().Set(CompactAndStrip(evt.GetPan().String()))
		output <- evt
	}

}

var CreatePadFunc = func(input, output chan Palette) {
	padChan := CreateNumberPump(16, 100)
	for evt := range input {
		evt.Log("Create pad...")
		evt.GetPad().Set(<-padChan)
		output <- evt
	}

}

var EncipherFunc = func(input, output chan Palette) {
	for evt := range input {
		evt.Log("Encipher the pan...")
		Encipher(evt)
		output <- evt
	}

}

var DecipherFunc = func(input, output chan Palette) {
	for evt := range input {
		evt.Log("Decipher the pan...")
		Decipher(evt)
		output <- evt
	}

}

func TokenFunc(binLength, lastLength int) PaletteFunc {
	fillStrChan := CreateNumberPump(6, 100)
	pf := func(input, output chan Palette) {
		for evt := range input {
			evt.Log("Create token...")

			pan := evt.GetPan().String()
			bin := Bin(pan, binLength)
			last := Last(pan, lastLength)
			fill := <-fillStrChan
			evt.GetToken().Set(bin + fill + last)
			//Loop to get an non-PAN version
			//No do/while ...
			for LuhnCheck(evt.GetToken().String()) {
				evt.Log("Illegal token created. Retry.")
				fill := <-fillStrChan
				evt.GetToken().Set(bin + fill + last)

			}
			output <- evt

		}
	}
	return pf
}
