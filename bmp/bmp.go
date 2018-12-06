package bmp

import (
	"io"
	"log"

	"github.com/chinx/morph"
	"golang.org/x/image/bmp"
)

func init() {
	morph.RegisterDecoder("bmp", Decode)
}

func Decode(r io.Reader) (d morph.Drawer, err error) {
	img, err := bmp.Decode(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}
	return morph.NewImage(img), nil
}
