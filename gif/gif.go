package gif

import (
	"image"
	"image/draw"
	"image/gif"
	"io"
	"log"

	"github.com/chinx/morph/internal"
)

func init() {
	internal.RegisterDecoder("gif", Decode)
	internal.RegisterEncoder("gif", Encode)
}

type Gif struct {
	img   *gif.GIF
	rst   *gif.GIF
	scale float64
	index int
}

func (g *Gif) Size(widthOrScale float64) {
	scale := widthOrScale
	if widthOrScale > 1 {
		width := float64(g.img.Config.Width)
		if width < scale {
			return
		}
		scale = scale / width
	}
	if scale == 1 {
		return
	}
	g.scale = scale
}

func (g *Gif) Encode(w io.Writer) error {
	result := g.img
	if g.index == len(g.img.Image) {
		result = g.rst
	}
	return gif.EncodeAll(w, result)
}

func (g *Gif) Next() bool {
	return g.index < len(g.img.Image)
}

func (g *Gif) Reset() {
	g.index = 0
}

func (g *Gif) Format() string {
	return "gif"
}

func (g *Gif) Image() image.Image {
	return g.img.Image[g.index]
}

func (g *Gif) Drawable() draw.Image {
	img := g.img.Image[g.index]
	rect := internal.ScaleRect(img.Bounds(), g.scale)
	drawer := image.NewPaletted(rect, img.Palette)
	g.rst.Image[g.index] = drawer
	g.index++
	return drawer
}

func Decode(r io.Reader) (d internal.Drawer, err error) {
	img, err := gif.DecodeAll(r)
	if err != nil {
		log.Println("failed to decode bmp reader: ", err)
		return nil, err
	}

	anim := &gif.GIF{
		BackgroundIndex: img.BackgroundIndex,
		Delay:           img.Delay,
		LoopCount:       img.LoopCount,
		Image:           make([]*image.Paletted, len(img.Image)),
	}

	return &Gif{img: img, rst: anim, scale: 1}, nil
}

func Encode(w io.Writer, img internal.Drawer) error {
	return gif.EncodeAll(w, img.(*Gif).rst)
}
