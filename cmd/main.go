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
	drawBlackAndWhiteFromTopInterval := flag.Int("draw-black-and-white-from-top-interval", 15,
		"in ms; the interval to sleep after drawing a B&W row")
	flag.Parse()

	image, err := asciianimator.NewImage(*imagePath)
	if err != nil {
		log.Fatalf("failed to parse the image image: %v", err)
	}

	image.DrawFromLeft(time.Duration(*drawFromLeftInterval) * time.Millisecond)
	time.Sleep(time.Duration(*leftToBlackAndWhiteInterval) * time.Millisecond)
	image.DrawBlackAndWhiteFromTop(time.Duration(*drawBlackAndWhiteFromTopInterval) * time.Millisecond)
}
