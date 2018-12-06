package internal

import (
	"bytes"
	"errors"
	"image"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	UnsupportedInput  = errors.New("unsupported input type")
	UnsupportedFormat = errors.New("unsupported image format")
)

var decoders = make(map[string]decoder)

type decoder func(io.Reader) (Drawer, error)

func RegisterDecoder(format string, dec decoder) {
	if _, ok := decoders[format]; ok {
		log.Println("format of decoder is exist: ", format)
	}
	decoders[format] = dec
}

func Decode(input interface{}) (drawer Drawer, err error) {
	switch t := input.(type) {
	case Drawer:
		drawer = t
		drawer.Reset()
	case io.ReadSeeker:
		drawer, err = decode(t)
	case io.Reader:
		data, err := ioutil.ReadAll(t)
		if err == nil {
			drawer, err = decode(bytes.NewReader(data))
		}
	case []byte:
		drawer, err = decode(bytes.NewReader(t))
	case string:
		file, err := os.Open(t)
		if err == nil {
			drawer, err = decode(file)
			file.Close()
		}
	default:
		err = UnsupportedInput
	}

	if err != nil {
		log.Println("failed to decode image: ", err)
	}
	return
}

func decode(rs io.ReadSeeker) (Drawer, error) {
	_, format, err := image.DecodeConfig(rs)
	if err != nil {
		log.Println(UnsupportedFormat, err)
		return nil, UnsupportedFormat
	}

	dec, ok := decoders[format]
	if !ok {
		log.Println(UnsupportedFormat, format)
		return nil, UnsupportedFormat
	}
	rs.Seek(0, 0)
	return dec(rs)
}
