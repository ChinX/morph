package bmp

import (
	"image/png"
	"io"
	"log"

	"github.com/chinx/morph/internal"
)

func init() {
	internal.RegisterDecoder("png", Decode)
	internal.RegisterEncoder("png", Encode)
}

func Decode(r io.Reader) (d internal.Drawer, err error) {
	img, err := png.Decode(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}
	return internal.NewImage(img, "png"), nil
}

func Encode(w io.Writer, img internal.Drawer) error {
	return png.Encode(w, img.Drawable())
}
