package main

import (
	"fmt"
	"os"

	"github.com/ploMP4/chafa-go"
)

const (
	FONT_WIDTH  = 18
	FONT_HEIGHT = 46
	N_CHANNELS  = 4
)

func main() {
	pixels, width, height, err := chafa.Load("./examples/image.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	config := chafa.CanvasConfigNew()
	defer chafa.CanvasConfigUnref(config)

	chafa.CanvasConfigSetGeometry(config, 40, 15)
	chafa.CanvasConfigSetCellGeometry(config, FONT_WIDTH, FONT_HEIGHT)
	chafa.CanvasConfigSetPixelMode(config, chafa.CHAFA_PIXEL_MODE_KITTY)
	chafa.CanvasConfigSetPassthrough(config, chafa.CHAFA_PASSTHROUGH_TMUX)

	widthNew := config.Width
	heightNew := config.Height

	chafa.CalcCanvasGeometry(
		width, height,
		&widthNew, &heightNew,
		float32(FONT_WIDTH)/float32(FONT_HEIGHT),
		false, false,
	)
	chafa.CanvasConfigSetGeometry(config, widthNew, heightNew)

	canvas := chafa.CanvasNew(config)
	defer chafa.CanvasUnRef(canvas)

	frame := chafa.FrameNew(
		pixels,
		chafa.CHAFA_PIXEL_RGBA8_UNASSOCIATED,
		width,
		height,
		width*N_CHANNELS,
	)

	img := chafa.ImageNew()
	chafa.ImageSetFrame(img, frame)

	placement := chafa.PlacementNew(img, 1)
	chafa.CanvasSetPlacement(canvas, placement)

	fmt.Println(chafa.CanvasPrint(canvas, nil))
}
