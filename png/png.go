package bmp

import (
	"image/png"
	"io"
	"log"

	"github.com/chinx/morph"
)

func init() {
	morph.RegisterDecoder("png", Decode)
}

func Decode(r io.Reader) (d morph.Drawer, err error) {
	img, err := png.Decode(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}
	return morph.NewImage(img), nil
}
