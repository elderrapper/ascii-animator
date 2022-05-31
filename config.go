package asciianimator

import (
	"fmt"
	"os"
	"time"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v3"
)

type Config struct {
	// ImagePath is the path to the image what will be animated.
	ImagePath string `yaml:"image_path" default:"image.ans"`

	// WithinStageInterval contains the time intervals to sleep between each action within the same stage.
	WithinStageInterval struct {
		// DrawFromLeft denotes the time interval to sleep between drawing each column
		// when the image is drawn from left to right.
		DrawFromLeft time.Duration `yaml:"draw_from_left" default:"15ms"`

		// DrawBlackAndWhiteFromTop denotes the time interval to sleep between drawing each row
		// when the B&W version of the image is drawn from top to bottom.
		DrawBlackAndWhiteFromTop time.Duration `yaml:"draw_black_and_white_from_top" default:"60ms"`

		// Sink denotes the time interval to sleep between each row sinks.
		Sink time.Duration `yaml:"sink" default:"120ms"`

		// Random denotes the time interval to sleep before displaying the next randomized image.
		Random time.Duration `yaml:"random" default:"750ms"`
	} `yaml:"within_stage_interval"`

	// BetweenStagesInterval contains the time interval to sleep between adjacent stages.
	BetweenStagesInterval struct {
		FromLeft2BlackAndWhiteFromTop time.Duration `yaml:"from_left_2_black_and_white_from_top" default:"1s"`
		BlackAndWhiteFromTop2Sink     time.Duration `yaml:"black_and_white_from_top_2_sink" default:"1s"`
		Sink2Random                   time.Duration `yaml:"sink_2_random" default:"1s"`
	} `yaml:"between_stage_interval"`

	// Random contains the configuration regarding how randomized images are generated.
	Random struct {
		// DrawProb denotes the probability to draw a randomized pixel.
		// The opposite of drawing a randomized pixel is to fill that pixel with the background color
		// so that the pixel seems to contain nothing to the audience.
		// In other words, the higher the probability is, the denser the generated image looks.
		DrawProb float64 `yaml:"draw_prob" default:"0.04"`

		// CharRange indicates the ASCII range of the possible characters.
		// Both boundaries are inclusive.
		//
		// Regarding the default values,
		// please check https://www.cs.mcgill.ca/~rwest/wikispeedia/wpcd/wp/a/ASCII.htm.
		CharRange struct {
			Start int `yaml:"start" default:"33"`
			End   int `yaml:"end" default:"126"`
		} `yaml:"char_range"`
	}
}

// NewConfig parses the file at confPath into a Config instance.
func NewConfig(confPath string) (*Config, error) {
	conf := &Config{}

	err := defaults.Set(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to set default values for the config: %w", err)
	}

	bs, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %w", err)
	}

	err = yaml.Unmarshal(bs, conf)
	if err != nil {
		return nil, fmt.Errorf("failed to YAML-unmarshal the file: %w", err)
	}

	return conf, nil
}
