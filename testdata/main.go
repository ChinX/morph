package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/chinx/morph"
	_ "github.com/chinx/morph/bmp"
	_ "github.com/chinx/morph/gif"
	_ "github.com/chinx/morph/jpeg"
	_ "github.com/chinx/morph/png"
	_ "github.com/chinx/morph/tiff"
	_ "github.com/chinx/morph/webp"
)

func main() {
	byteData, err := ioutil.ReadFile("testdata/branches.jpg")
	if err != nil {
		log.Fatalf("failed to read image: %v", err)
	}

	draw, err := morph.DecodeBytes(byteData)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}

	err = morph.Scale(draw, 200)
	if err != nil {
		log.Fatalf("failed scale for image: %v", err)
	}

	file, _ := os.Create("testdata/branches_out.jpg")
	draw.Encode(file)
	file.Close()

	byteData, err = ioutil.ReadFile("testdata/branches.png")
	if err != nil {
		log.Fatalf("failed to read image: %v", err)
	}

	draw, err = morph.DecodeBytes(byteData)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}

	err = morph.Scale(draw, 200)
	if err != nil {
		log.Fatalf("failed scale for image: %v", err)
	}

	file, _ = os.Create("testdata/branches_out.png")

	draw.Encode(file)
	file.Close()

	byteData, err = ioutil.ReadFile("testdata/timg.gif")
	if err != nil {
		log.Fatalf("failed to read image: %v", err)
	}

	draw, err = morph.DecodeBytes(byteData)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}

	err = morph.Scale(draw, 200)
	if err != nil {
		log.Fatalf("failed scale for image: %v", err)
	}

	file, _ = os.Create("testdata/timg_out.gif")

	draw.Encode(file)
	file.Close()
}
