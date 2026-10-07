package auximageprovider

import (
	"github.com/imgproxy/imgproxy/v4/ensure"
	"github.com/imgproxy/imgproxy/v4/env"
)

var (
	DIMGPROXY_WATERMARK_DATA = env.String("IMGPROXY_WATERMARK_DATA")
	DIMGPROXY_WATERMARK_PATH = env.ExistingFilePath("IMGPROXY_WATERMARK_PATH")
	DIMGPROXY_WATERMARK_URL  = env.String("IMGPROXY_WATERMARK_URL")

	DIMGPROXY_FALLBACK_IMAGE_DATA = env.String("IMGPROXY_FALLBACK_IMAGE_DATA")
	DIMGPROXY_FALLBACK_IMAGE_PATH = env.ExistingFilePath("IMGPROXY_FALLBACK_IMAGE_PATH")
	DIMGPROXY_FALLBACK_IMAGE_URL  = env.String("IMGPROXY_FALLBACK_IMAGE_URL")
)

// DynamicConfig holds the configuration for the auxiliary image provider
type DynamicConfig struct {
	Base64Data string
	Path       string
	URL        string
}

// NewDefaultDynamicConfig creates a new default configuration for the auxiliary image provider
func NewDefaultDynamicConfig() DynamicConfig {
	return DynamicConfig{
		Base64Data: "",
		Path:       "",
		URL:        "",
	}
}

// LoadWatermarkDynamicConfigFromEnv loads the watermark configuration from the environment
func LoadWatermarkDynamincConfigFromEnv(c *DynamicConfig) (*DynamicConfig, error) {
	c = ensure.Ensure(c, NewDefaultDynamicConfig)

	DIMGPROXY_WATERMARK_DATA.Parse(&c.Base64Data)
	DIMGPROXY_WATERMARK_PATH.Parse(&c.Path)
	DIMGPROXY_WATERMARK_URL.Parse(&c.URL)

	return c, nil
}

// LoadFallbackDynamicConfigFromEnv loads the fallback configuration from the environment
func LoadFallbackdynaminConfigFromEnv(c *DynamicConfig) (*DynamicConfig, error) {
	c = ensure.Ensure(c, NewDefaultDynamicConfig)

	DIMGPROXY_FALLBACK_IMAGE_DATA.Parse(&c.Base64Data)
	DIMGPROXY_FALLBACK_IMAGE_PATH.Parse(&c.Path)
	DIMGPROXY_FALLBACK_IMAGE_URL.Parse(&c.URL)

	return c, nil
}
