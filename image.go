package asciianimator

import (
	"fmt"
	"os"
	"time"

	"github.com/leaanthony/go-ansi-parser"
)

// Image is an 2D ASCII artwork.
type Image [][]*ansi.StyledText

// NewImage parses an image that is represented by ASCII art and stored in an .ans file.
func NewImage(imagePath string) (Image, error) {
	bs, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the image file: %w", err)
	}

	var image [][]*ansi.StyledText
	str := string(bs)
	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			// I'm not sure why .ans files downloaded from Ascii Art converter [1]
			// start with a seemly invalid ANSI escape sequence: `^[[107;40m`.
			//
			// [1] https://manytools.org/hacker-tools/convert-images-to-ascii-art/
			row, err := ansi.Parse(
				str[start:i],
				ansi.WithIgnoreInvalidCodes(),
			)
			if err != nil {
				return nil, fmt.Errorf("failed to parse the read image: %w", err)
			}

			image = append(image, row)
			start = i + 1
		}
	}
	return image, nil
}

// DrawFromLeft draws the image from left and sleeps sleepInterval after drawing a column.
// The cursor position remains unchanged after this function returns.
func (img Image) DrawFromLeft(sleepInterval time.Duration) {
	for j := range img[0] {
		for i := range img {
			fmt.Print(img[i][j].String())
			moveDown(1)
			moveLeft(1)
		}

		moveUp(len(img))
		moveRight(1)
		time.Sleep(sleepInterval)
	}
	moveLeft(len(img[0]))
}

