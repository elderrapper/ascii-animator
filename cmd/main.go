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
	drawBlackAndWhiteFromTopInterval := flag.Int("draw-black-and-white-from-top-interval", 45,
		"in ms; the interval to sleep after drawing a B&W row")
	blackAndWhiteToSinkInterval := flag.Int("black-and-white-to-sink-interval", 2000,
		"in ms; the interval to sleep between drawing the B&W version of the image from top and make it sink")
	sinkInterval := flag.Int("sink-interval", 75,
		"in ms; the interval to sleep when sinking each row of the image")
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
}
