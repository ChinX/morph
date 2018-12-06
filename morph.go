package morph

import (
	"bytes"
	"errors"
	"image"
	"image/draw"
	"io"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/graphics-go/graphics"
)

func Blur(drawer Drawer, opt *graphics.BlurOptions) error {
	return scan(drawer, func(dst draw.Image, src image.Image) error {
		return graphics.Blur(dst, src, opt)
	}, 0)
}

func Rotate(drawer Drawer, opt *graphics.RotateOptions) error {
	return scan(drawer, func(dst draw.Image, src image.Image) error {
		return graphics.Rotate(dst, src, opt)
	}, 0)
}

func Scale(drawer Drawer, widthOrScale float64) error {
	return scan(drawer, func(dst draw.Image, src image.Image) error {
		return graphics.Scale(dst, src)
	}, widthOrScale)
}

func Thumbnail(drawer Drawer, widthOrScale float64) error {
	return scan(drawer, func(dst draw.Image, src image.Image) error {
		return graphics.Thumbnail(dst, src)
	},widthOrScale)
}

func scan(drawer Drawer, fn func(dst draw.Image, src image.Image) error, widthOrScale float64) error {
	drawer.Reset()
	drawer.Size(widthOrScale)
	for drawer.Next() {
		src := drawer.Image()
		dst := drawer.Drawable()
		if err := fn(dst, src); err != nil {
			return err
		}
	}
	return nil
}

func Decode(r io.Reader) (Drawer, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("failed to read image reader: ", err)
		return nil, err
	}
	return DecodeBytes(data)
}

func DecodeBytes(data []byte) (Drawer, error) {
	buffer := bytes.NewReader(data)
	_, format, err := image.DecodeConfig(buffer)
	if err != nil {
		log.Println("Unsupported image format: ", err)
		return nil, err
	}

	dec, ok := decoders[format]
	if !ok {
		err = errors.New("Unsupported image format: " + format)
		log.Println(err)
	}
	buffer.Seek(0, 0)
	return dec(buffer)
}
