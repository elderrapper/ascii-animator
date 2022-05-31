package main

import (
	"flag"
	"log"
	"time"

	asciianimator "github.com/davidhsingyuchen/ascii-animator"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	config, err := asciianimator.NewConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load the config file: %v", err)
	}

	image, err := asciianimator.NewImage(config.ImagePath)
	if err != nil {
		log.Fatalf("failed to parse the image: %v", err)
	}

	image.DrawFromLeft(config.WithinStageInterval.DrawFromLeft)
	time.Sleep(config.BetweenStagesInterval.FromLeft2BlackAndWhiteFromTop)
	image.DrawBlackAndWhiteFromTop(config.WithinStageInterval.DrawBlackAndWhiteFromTop)
	time.Sleep(config.BetweenStagesInterval.BlackAndWhiteFromTop2Sink)
	image.Sink(config.WithinStageInterval.Sink)
	time.Sleep(config.BetweenStagesInterval.Sink2Random)
	image.RandomizeColorAndChars(
		config.WithinStageInterval.Random,
		config.Random.DrawProb,
		config.Random.CharRange.Start,
		config.Random.CharRange.End,
	)
}
