package asciianimator

import (
	"fmt"
	"math/rand"
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
			row, err := ansi.Parse(str[start:i])
			if err != nil {
				return nil, fmt.Errorf("failed to parse the read image: %w", err)
			}

			image = append(image, row)
			start = i + 1
		}
	}
	return image, nil
}

func (img Image) Clone() Image {
	clone := make(Image, len(img))
	for i := range img {
		clone[i] = make([]*ansi.StyledText, len(img[i]))
		copy(clone[i], img[i])
	}
	return clone
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

// DrawBlackAndWhiteFromTop draws a B&W version of the image from top.
// The cursor position remains unchanged after this function returns.
func (img Image) DrawBlackAndWhiteFromTop(sleepInterval time.Duration) {
	clone := img.Clone()
	clone.toBlackAndWhite()
	clone.drawFromTop(sleepInterval)
}

func (img Image) toBlackAndWhite() {
	for _, row := range img {
		for _, pixel := range row {
			if !isBlack(pixel.FgCol) {
				pixel.FgCol = ansi.Cols[colorWhite]
			}
		}
	}
}

func (img Image) drawFromTop(sleepInterval time.Duration) {
	for _, row := range img {
		for _, pixel := range row {
			fmt.Print(pixel.String())
		}

		fmt.Println()
		time.Sleep(sleepInterval)
	}
	moveUp(len(img))
}

// Sink make the drawn image sink.
// The cursor position remains unchanged after this function returns.
func (img Image) Sink(sleepInterval time.Duration) {
	for n := len(img) - 1; ; n-- {
		eraseEntireLine()
		moveDown(1)
		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			for _, pixel := range img[i] {
				fmt.Print(pixel.String())
			}
			fmt.Println()
		}

		moveUp(n)
		time.Sleep(sleepInterval)
	}
	moveUp(len(img))
}

// RandomizeColorAndText displays an image with the same dimension of the image,
// but both the colors and the chars are randomized.
// This function is designed to be run forever.
func (img Image) RandomizeColorAndChars(
	sleepInterval time.Duration, probToDraw float64, startChar, endChar int) {

	clone := img.Clone()
	rand.Seed(time.Now().UnixNano())
	moveUp(1)

	for {
		for _, row := range clone {
			for _, pixel := range row {
				pixel.Label = string(byte(rand.Intn(endChar-startChar+1) + startChar))
				if probToDraw > rand.Float64() {
					pixel.FgCol = ansi.Cols[rand.Intn(numColors)]
				} else {
					pixel.FgCol = ansi.Cols[colorBlack]
				}

				fmt.Print(pixel)
			}
			fmt.Println()
		}

		moveUp(len(clone))
		time.Sleep(sleepInterval)
	}
}
