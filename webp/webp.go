package bmp

import (
	"image/jpeg"
	"io"
	"log"

	"github.com/chinx/morph/internal"
	"golang.org/x/image/webp"
)

func init() {
	internal.RegisterDecoder("webp", Decode)
	internal.RegisterEncoder("webp", Encode)
}

func Decode(r io.Reader) (d internal.Drawer, err error) {
	img, err := webp.Decode(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}
	return internal.NewImage(img, "webp"), nil
}

func Encode(w io.Writer, img internal.Drawer) error {
	return jpeg.Encode(w, img.Drawable(), nil)
}
