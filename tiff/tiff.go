package bmp

import (
	"io"
	"log"

	"github.com/chinx/morph/internal"
	"golang.org/x/image/tiff"
)

func init() {
	internal.RegisterDecoder("tiff", Decode)
	internal.RegisterEncoder("tiff", Encode)
}

func Decode(r io.Reader) (d internal.Drawer, err error) {
	img, err := tiff.Decode(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}
	return internal.NewImage(img, "tiff"), nil
}

func Encode(w io.Writer, img internal.Drawer) error {
	return tiff.Encode(w, img.Drawable(), nil)
}
