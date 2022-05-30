package main

import (
	"flag"
	"fmt"
	"log"

	asciianimator "github.com/davidhsingyuchen/ascii-animator"
)

func main() {
	imagePath := flag.String("image-path", "image.ans", "path to the image")
	flag.Parse()

	image, err := asciianimator.NewImage(*imagePath)
	if err != nil {
		log.Fatalf("failed to parse the image image: %v", err)
	}

	for _, row := range image {
		for _, pixel := range row {
			fmt.Print(pixel.String())
		}
		fmt.Println()
	}
}
