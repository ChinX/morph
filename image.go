package morph

import (
	"image"
	"image/color"
	"image/draw"
	"io"
	"log"

	"golang.org/x/image/bmp"
)

var decoders = make(map[string]decoder)

type decoder func(io.Reader) (Drawer, error)

func RegisterDecoder(format string, dec decoder) {
	if _, ok := decoders[format]; ok {
		log.Println("format of decoder is exist: ", format)
	}
	decoders[format] = dec
}

type Drawer interface {
	Size(widthOrScale float64)
	Encode(w io.Writer) error
	Next() bool
	Reset()
	Image() image.Image
	Drawable() draw.Image
}

type Image struct {
	img     image.Image
	rst     draw.Image
	rect    image.Rectangle
	morphed bool
}

func NewImage(img image.Image) *Image {
	return &Image{img: img, rect: img.Bounds()}
}

func (i *Image) Size(widthOrScale float64) {
	if widthOrScale == 0 {
		return
	}
	scale := widthOrScale
	if widthOrScale > 1 {
		width := float64(i.rect.Dx())
		if width < scale {
			return
		}
		scale = scale / width
	}
	if scale == 1 {
		return
	}
	i.rect = ScaleRect(i.rect, scale)
}

func (i *Image) Encode(w io.Writer) error {
	result := i.img
	if i.rst != nil {
		result = i.rst
	}
	return bmp.Encode(w, result)
}

func (i *Image) Next() bool {
	return !i.morphed
}

func (i *Image) Reset() {
	i.morphed = false
}

func (i *Image) Image() image.Image {
	return i.img
}

func (i *Image) Drawable() draw.Image {
	if i.rst == nil {
		i.rst = NewDrawableSize(i.img, i.rect)
	}
	i.morphed = true
	return i.rst
}

func NewDrawableSize(p image.Image, r image.Rectangle) draw.Image {
	switch p := p.(type) {
	case *image.RGBA:
		return image.NewRGBA(r)
	case *image.RGBA64:
		return image.NewRGBA64(r)
	case *image.NRGBA:
		return image.NewNRGBA(r)
	case *image.NRGBA64:
		return image.NewNRGBA64(r)
	case *image.Alpha:
		return image.NewAlpha(r)
	case *image.Alpha16:
		return image.NewAlpha16(r)
	case *image.Gray:
		return image.NewGray(r)
	case *image.Gray16:
		return image.NewGray16(r)
	case *image.Paletted:
		pl := make(color.Palette, len(p.Palette))
		copy(pl, p.Palette)
		return image.NewPaletted(r, pl)
	case *image.CMYK:
		return image.NewCMYK(r)
	default:
		return image.NewRGBA(r)
	}
}

func ScaleRect(rect image.Rectangle, scale float64) image.Rectangle {
	rect.Min.X = int(float64(rect.Min.X)*scale + 0.5)
	rect.Min.Y = int(float64(rect.Min.Y)*scale + 0.5)
	rect.Max.X = int(float64(rect.Max.X)*scale + 0.5)
	rect.Max.Y = int(float64(rect.Max.Y)*scale + 0.5)
	return rect
}
