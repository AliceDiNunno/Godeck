package imagebuilder

import (
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"os"
)

func CreateImage(height int, width int) *image.RGBA {
	basicImage := image.NewRGBA(image.Rect(0, 0, width, height))
	backgroundColor := image.Transparent

	draw.Draw(basicImage, basicImage.Bounds(), backgroundColor, image.Point{}, draw.Src)

	return basicImage
}

func LoadPng(path string) *image.RGBA {
	existingImageFile, err := os.Open(path)
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "png"
	_, _, err = image.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)

	if err != nil {
		log.Errorln(err)
		return nil
	}

	return ImageToRGBA(loadedImage)
}

func ImageToRGBA(src image.Image) *image.RGBA {
	// No conversion needed if image is an *image.RGBA.
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}

	// Use the image/draw package to convert to *image.RGBA.
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

func ResizeImage(image *image.RGBA, width int, height int) *image.RGBA {
	m := resize.Resize(uint(width), uint(height), image, resize.Lanczos3)

	return ImageToRGBA(m)
}

func AddPngToBaseImage(baseImage *image.RGBA, png *image.RGBA, x int, y int) *image.RGBA {
	if png.Bounds().Max.X > baseImage.Bounds().Max.X || png.Bounds().Max.Y > baseImage.Bounds().Max.Y {
		png = ResizeImage(png, baseImage.Bounds().Max.X, baseImage.Bounds().Max.Y)
	}

	draw.Draw(baseImage, baseImage.Bounds(), png, image.Point{x, y}, draw.Src)

	return baseImage
}