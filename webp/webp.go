package bmp

import (
	"io"
	"log"

	"github.com/chinx/morph"
	"golang.org/x/image/webp"
)

func init() {
	morph.RegisterDecoder("webp", Decode)
}

func Decode(r io.Reader) (d morph.Drawer, err error) {
	img, err := webp.Decode(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}
	return morph.NewImage(img), nil
}
