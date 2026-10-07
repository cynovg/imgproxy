package auximageprovider

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