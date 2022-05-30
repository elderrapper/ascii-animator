package main

import (
	"flag"
	"log"
	"time"

	asciianimator "github.com/davidhsingyuchen/ascii-animator"
)

func main() {
	imagePath := flag.String("image-path", "image.ans", "path to the image")
	drawFromLeftInterval := flag.Int("draw-from-left-interval", 15,
		"in ms; the interval to sleep after drawing a column")
	leftToBlackAndWhiteInterval := flag.Int("left-to-black-and-white-interval", 2000,
		"in ms; the interval to sleep between drawing the image from left and drawing its B&W version from top")
	drawBlackAndWhiteFromTopInterval := flag.Int("draw-black-and-white-from-top-interval", 60,
		"in ms; the interval to sleep after drawing a B&W row")
	blackAndWhiteToSinkInterval := flag.Int("black-and-white-to-sink-interval", 2000,
		"in ms; the interval to sleep between drawing the B&W version of the image from top and make it sink")
	sinkInterval := flag.Int("sink-interval", 120,
		"in ms; the interval to sleep when sinking each row of the image")
	sinkToRandomizeColorAndCharsInterval := flag.Int("sink-to-randomize-color-and-chars-interval", 2000,
		"in ms; the interval to sleep between sinking each row of the image and randomizing "+
			"both the color and the Chars of it")
	randomizeColorAndCharsInterval := flag.Int("randomize-color-and-chars-interval", 750,
		"in ms; the interval to sleep when randomizing both the color and the chars of the image")

	randomizeColorAndCharsDrawProb := flag.Float64("randomize-color-and-chars-draw-prob", 0.04,
		"the probability to draw a pixel; the higher it is, the denser the generated image looks")
	// Printable ASCII characters are from 33 to 126.
	// Ref. https://www.cs.mcgill.ca/~rwest/wikispeedia/wpcd/wp/a/ASCII.htm
	randomizeCharsRangeStart := flag.Int("randomize-chars-range-start", 33,
		"the start of the ASCII range of the possible characters")
	randomizeCharsRangeEnd := flag.Int("randomize-chars-range-end", 126,
		"the end of the ASCII range of the possible characters")
	flag.Parse()

	image, err := asciianimator.NewImage(*imagePath)
	if err != nil {
		log.Fatalf("failed to parse the image image: %v", err)
	}

	image.DrawFromLeft(time.Duration(*drawFromLeftInterval) * time.Millisecond)
	time.Sleep(time.Duration(*leftToBlackAndWhiteInterval) * time.Millisecond)
	image.DrawBlackAndWhiteFromTop(time.Duration(*drawBlackAndWhiteFromTopInterval) * time.Millisecond)
	time.Sleep(time.Duration(*blackAndWhiteToSinkInterval) * time.Millisecond)
	image.Sink(time.Duration(*sinkInterval) * time.Millisecond)
	time.Sleep(time.Duration(*sinkToRandomizeColorAndCharsInterval) * time.Millisecond)
	image.RandomizeColorAndChars(
		time.Duration(*randomizeColorAndCharsInterval)*time.Millisecond,
		*randomizeColorAndCharsDrawProb,
		*randomizeCharsRangeStart,
		*randomizeCharsRangeEnd,
	)
}
