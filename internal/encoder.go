package internal

import (
	"io"
	"log"
	"os"
	"strings"
)

var encoders = make(map[string]encoder)

type encoder func(io.Writer, Drawer) error

func RegisterEncoder(format string, enc encoder) {
	if _, ok := encoders[format]; ok {
		log.Println("format of encoder is exist: ", format)
	}
	encoders[format] = enc
}

func Encode(output interface{}, img Drawer) (filename string, err error) {
	switch t := output.(type) {
	case io.Writer:
		err = encode(t, img)
	case string:
		if strings.HasSuffix(t, img.Format()){
			filename = t
		}else{
			filename = t + "."+img.Format()
		}
		var file *os.File
		if file, err = os.Create(filename); err == nil {
			err = encode(file, img)
			file.Close()
		}
	default:
		err = UnsupportedInput
	}

	if err != nil {
		log.Println("failed to encode image: ", err)
	}
	return
}

func encode(w io.Writer, img Drawer) error {
	enc, ok := encoders[img.Format()]
	if !ok {
		log.Println(UnsupportedFormat, img.Format())
		return UnsupportedFormat
	}
	return enc(w, img)
}
