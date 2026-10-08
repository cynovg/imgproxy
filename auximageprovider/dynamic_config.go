package auximageprovider

import (
	"github.com/imgproxy/imgproxy/v4/ensure"
	"github.com/imgproxy/imgproxy/v4/env"
)

var (
	IMGPROXY_WATERMARK_FONT_PATH = env.String("IMGPROXY_WATERMARK_FONT_PATH")
	IMGPROXY_WATERMARK_FONT_SIZE = env.Float("IMGPROXY_WATERMARK_FONT_SIZE")
)

// DynamicConfig holds the configuration for the auxiliary image provider
type DynamicConfig struct {
	Path       string
	Size       float64
}

// NewDefaultDynamicConfig creates a new default configuration for the auxiliary image provider
func NewDefaultDynamicConfig() DynamicConfig {
	return DynamicConfig{
		Path:       "",
		Size:        0,
	}
}

func LoadWatermarkDynamincConfigFromEnv(c *DynamicConfig) (*DynamicConfig, error) {
	c = ensure.Ensure(c, NewDefaultDynamicConfig)

	IMGPROXY_WATERMARK_FONT_PATH.Parse(&c.Path)
	IMGPROXY_WATERMARK_FONT_SIZE.Parse(&c.Size)

	return c, nil
}
