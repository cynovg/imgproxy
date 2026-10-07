package auximageprovider

import (
	"bytes"
	"context"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"net/http"

	"github.com/fogleman/gg"
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
		fontPath := "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"
		fontSize := float64(22)
		bytes, err := GenerateWatermarkPNG(text, fontSize, fontPath)
		if err != nil {
			return nil, nil, err
		}

		data, err := s.idf.NewFromBytes(bytes)
		if err != nil {
			return nil, nil, err
		}
		return data, make(http.Header), nil
	}
	return nil, s.headers.Clone(), nil
}

func GenerateWatermarkPNG(text string, fontSize float64, fontPath string) ([]byte, error) {
	measureDC := gg.NewContext(1, 1)

	if err := measureDC.LoadFontFace(fontPath, fontSize); err != nil {
		return nil, err
	}

	textWidth, textHeight := measureDC.MeasureString(text)

	const padding = 4.0
	width := int(math.Ceil(textWidth)) + int(padding*2)
	height := int(math.Ceil(textHeight)) + int(padding*2)

	dc := gg.NewContext(width, height)

	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		return nil, err
	}

	dc.SetColor(color.RGBA{255, 255, 255, 255})
	dc.DrawStringAnchored(text, float64(width)/2, float64(height)/2, 0.5, 0.5)

	var buf bytes.Buffer
	img := dc.Image()

	dst := image.NewNRGBA(img.Bounds())
	draw.Draw(dst, dst.Bounds(), img, img.Bounds().Min, draw.Src)

	if err := png.Encode(&buf, dst); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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
