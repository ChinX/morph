package morph

import (
	"errors"
	"image"
	"image/draw"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/chinx/morph/internal"
)

var Unsupported = errors.New("unsupported input type")

func Blur(input, output interface{}, opt *graphics.BlurOptions) (filename string, err error) {
	return scan(input, output, func(dst draw.Image, src image.Image) error {
		return graphics.Blur(dst, src, opt)
	}, 0)
}

func Rotate(input, output interface{}, opt *graphics.RotateOptions) (filename string, err error) {
	return scan(input, output, func(dst draw.Image, src image.Image) error {
		return graphics.Rotate(dst, src, opt)
	}, 0)
}

func Scale(input, output interface{}, widthOrScale float64) (filename string, err error) {
	return scan(input, output, func(dst draw.Image, src image.Image) error {
		return graphics.Scale(dst, src)
	}, widthOrScale)
}

func Thumbnail(input, output interface{}, widthOrScale float64) (filename string, err error) {
	return scan(input, output, func(dst draw.Image, src image.Image) error {
		return graphics.Thumbnail(dst, src)
	}, widthOrScale)
}

func scan(input, output interface{}, fn func(dst draw.Image, src image.Image) error, widthOrScale float64) (filename string, err error) {
	drawer, err := internal.Decode(input)
	if err != nil {
		return "", err
	}
	drawer.Size(widthOrScale)
	for drawer.Next() {
		src := drawer.Image()
		dst := drawer.Drawable()
		if err := fn(dst, src); err != nil {
			return "", err
		}
	}
	return internal.Encode(output, drawer)
}
