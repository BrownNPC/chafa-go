package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/ploMP4/chafa-go"
)

const (
	FONT_WIDTH  = 11
	FONT_HEIGHT = 24
	N_CHANNELS  = 4
)

func main() {
	pixels, width, height, err := loadPNGToRGBA("./cmd/image.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	config := chafa.CanvasConfigNew()
	defer chafa.CanvasConfigUnref(config)

	chafa.CanvasConfigSetGeometry(config, 30, 30)
	chafa.CanvasConfigSetCellGeometry(config, FONT_WIDTH, FONT_HEIGHT)

	var wNew int32 = config.Width
	var hNew int32 = config.Height

	chafa.CalcCanvasGeometry(
		width,
		height,
		&wNew,
		&hNew,
		float32(FONT_WIDTH)/float32(FONT_HEIGHT),
		false,
		false,
	)
	chafa.CanvasConfigSetGeometry(config, wNew, hNew)

	// Create canvas
	canvas := chafa.CanvasNew(config)
	defer chafa.CanvasUnRef(canvas)

	// Draw pixels to canvas
	chafa.CanvasDrawAllPixels(
		canvas,
		chafa.CHAFA_PIXEL_RGBA8_UNASSOCIATED,
		pixels,
		int32(width),
		int32(height),
		int32(width)*N_CHANNELS,
	)

	gs := chafa.CanvasPrint(canvas, nil)

	fmt.Println(gs)
}

func loadPNGToRGBA(path string) ([]uint8, int32, int32, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, 0, 0, err
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Ensure it's in RGBA format
	rgbaImg := image.NewRGBA(bounds)
	draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)

	return rgbaImg.Pix, int32(width), int32(height), nil
}
