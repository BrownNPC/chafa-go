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
	N_CHANNELS = 4
)

func main() {
	pixels, width, height, err := loadPNGToRGBA("./cmd/image.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Specify the symbols we want
	symbolMap := chafa.SymbolMapNew()
	defer chafa.SymbolMapUnref(symbolMap)

	chafa.SymbolMapAddByTags(symbolMap, chafa.CHAFA_SYMBOL_TAG_ALL)

	//  Set up a configuration with the symbols and the canvas size in characters
	config := chafa.CanvasConfigNew()
	defer chafa.CanvasConfigUnref(config)

	chafa.CanvasConfigSetGeometry(config, 40, 15)
	chafa.CanvasConfigSetSymbolMap(config, symbolMap)
	// chafa.CanvasConfigSetPixelMode(config, chafa.CHAFA_PIXEL_MODE_KITTY)

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

func loadPNGToRGBA(path string) ([]uint8, int, int, error) {
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

	return rgbaImg.Pix, width, height, nil
}
