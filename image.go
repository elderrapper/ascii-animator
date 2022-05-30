package asciianimator

import (
	"fmt"
	"os"

	"github.com/leaanthony/go-ansi-parser"
)

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
