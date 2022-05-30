package main

import (
	"flag"
	"log"
	"time"

	asciianimator "github.com/davidhsingyuchen/ascii-animator"
)

func main() {
	imagePath := flag.String("image-path", "image.ans", "path to the image")
	drawFromLeftInterval := flag.Int("draw-from-left-interval", 50,
		"in ms; the interval to sleep after drawing a column")
	flag.Parse()

	image, err := asciianimator.NewImage(*imagePath)
	if err != nil {
		log.Fatalf("failed to parse the image image: %v", err)
	}

	image.DrawFromLeft(time.Duration(*drawFromLeftInterval) * time.Millisecond)
}
