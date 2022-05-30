package asciianimator

import "github.com/leaanthony/go-ansi-parser"

const (
	colorBlack = 0
	colorWhite = 15

	// https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
	colorBlackGrayscaleStart = 232
	colorBlackGrayscaleEnd   = 243

	numColors = 256
)

func isBlack(color *ansi.Col) bool {
	return color.Id == colorBlack ||
		(color.Id >= colorBlackGrayscaleStart && color.Id <= colorBlackGrayscaleEnd)
}
