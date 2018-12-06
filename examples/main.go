package main

import (
	"io/ioutil"
	"log"

	"github.com/chinx/morph"
	_ "github.com/chinx/morph/bmp"
	_ "github.com/chinx/morph/gif"
	_ "github.com/chinx/morph/jpeg"
	_ "github.com/chinx/morph/png"
	_ "github.com/chinx/morph/tiff"
	_ "github.com/chinx/morph/webp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	byteData, err := ioutil.ReadFile("testdata/branches.jpg")
	if err != nil {
		log.Fatalf("failed to read image: %v", err)
	}

	_, err = morph.Scale(byteData, "testdata/branches_out.jpg", 200)
	if err != nil {
		log.Fatalf("failed scale for image: %v", err)
	}

	byteData, err = ioutil.ReadFile("testdata/branches.png")
	if err != nil {
		log.Fatalf("failed to read image: %v", err)
	}

	_, err = morph.Scale(byteData, "testdata/branches_out.png", 200)
	if err != nil {
		log.Fatalf("failed scale for image: %v", err)
	}

	byteData, err = ioutil.ReadFile("testdata/timg.gif")
	if err != nil {
		log.Fatalf("failed to read image: %v", err)
	}

	_, err = morph.Scale(byteData, "testdata/timg_out.gif", 200)
	if err != nil {
		log.Fatalf("failed scale for image: %v", err)
	}

}
