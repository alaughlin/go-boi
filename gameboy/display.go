package gameboy

type display struct {
	width      int
	height     int
	ScreenData []byte
}

func initalizeDisplay(width int, height int) *display {
	screenData := make([]byte, width*height*4)

	for i := range screenData {
		screenData[i] = 255
	}

	return &display{
		width:      width,
		height:     height,
		ScreenData: screenData,
	}
}
