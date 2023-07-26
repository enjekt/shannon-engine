package shannon_engine

type paletteFunc = func(input, output chan *palette)

var CompactAndStripPanFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Compact and strip..")
		evt.Pan = Pan(CompactAndStrip(string(evt.Pan)))
		output <- evt
	}

}

var CreatePadFunc = func(input, output chan *palette) {
	rndNumberStrChan := CreateNumberPump(16, 100)
	for evt := range input {
		evt.Log("Create Pad...")
		evt.Pad = Pad(CompactAndStrip(string(<-rndNumberStrChan)))
		output <- evt
	}

}

var PadPanFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Pad the Pan...")
		evt.PaddedPan = PaddedPan(Encipher(evt.Pan, evt.Pad))
		output <- evt
	}

}

var CreateTokenFunc = func(input, output chan *palette) {
	for evt := range input {
		evt.Log("Create Token...")
		//TODO
		output <- evt
	}
}
