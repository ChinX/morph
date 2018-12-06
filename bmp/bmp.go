package bmp

import (
	"io"
	"log"

	"github.com/chinx/morph/internal"
	"golang.org/x/image/bmp"
)

func init() {
	internal.RegisterDecoder("bmp", Decode)
	internal.RegisterEncoder("bmp", Encode)
}

func Decode(r io.Reader) (d internal.Drawer, err error) {
	img, err := bmp.Decode(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}
	return internal.NewImage(img, "bmp"), nil
}

func Encode(w io.Writer, img internal.Drawer) error {
	return bmp.Encode(w, img.Drawable())
}
