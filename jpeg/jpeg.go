package bmp

import (
	"image/jpeg"
	"io"
	"log"

	"github.com/chinx/morph/internal"
)

func init() {
	internal.RegisterDecoder("jpeg", Decode)
	internal.RegisterEncoder("jpeg", Encode)
}

func Decode(r io.Reader) (d internal.Drawer, err error) {
	img, err := jpeg.Decode(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}
	return internal.NewImage(img, "jpeg"), nil
}

func Encode(w io.Writer, img internal.Drawer) error {
	return jpeg.Encode(w, img.Drawable(), nil)
}
