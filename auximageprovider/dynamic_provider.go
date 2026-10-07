package auximageprovider

import (
	"context"
	"net/http"
	"encoding/base64"

	"github.com/imgproxy/imgproxy/v4/imagedata"
	"github.com/imgproxy/imgproxy/v4/options"
	"github.com/imgproxy/imgproxy/v4/options/keys"
)

// staticProvider is a simple implementation of ImageProvider, which returns
// a static saved image data and headers.
type dynamicProvider struct {
	data    imagedata.ImageData
	headers http.Header
	idf     imagedata.Factory
}

// Get returns the static image data and headers stored in the provider.
func (s *dynamicProvider) Get(_ context.Context, o *options.Options) (imagedata.ImageData, http.Header, error) {
	if text := o.GetString(keys.WatermarkText, ""); text != "" {
		bytes, err := base64.RawURLEncoding.DecodeString(text)
		if err != nil {
			return nil,nil, err
		}

		data, err := s.idf.NewFromBytes(bytes)
		if err != nil {
			return nil, nil, err
		}
		return data, make(http.Header), nil
	}
	return s.data.Ref(), s.headers.Clone(), nil
}

// Close releases the static image data held by the provider.
func (s *dynamicProvider) Close() error {
	if s.data != nil {
		return s.data.Close()
	}
	return nil
}

// NewDynamicProvider creates a new ImageProvider from either a base64 string, file path, or URL
func NewDynamicProvider(
	ctx context.Context,
	c *DynamicConfig,
	desc string,
	idf imagedata.Factory,
) (Provider, error) {
	var (
		data    imagedata.ImageData
		headers = make(http.Header)
		err     error
	)

	if err != nil {
		return nil, err
	}

	return &dynamicProvider{
		data:    data,
		headers: headers,
		idf:     idf,
	}, nil
}
