package main

import (
	"fmt"
	"os"

	"github.com/ploMP4/chafa-go"
	"golang.org/x/sys/unix"
)

const (
	PIX_WIDTH  = 3
	PIX_HEIGHT = 3
	N_CHANNELS = 4
)

type TermSize struct {
	widthCells   int32
	heightCells  int32
	widthPixels  int32
	heightPixels int32
}

func detectTerminal() (*chafa.TermInfo, chafa.CanvasMode, chafa.PixelMode, chafa.Passthrough, *chafa.SymbolMap) {
	//Examine the environment variables and guess what the terminal can do
	termInfo := chafa.TermDbDetect(chafa.TermDbGetDefault(), os.Environ())

	// Pick the most high-quality rendering possible
	mode := chafa.TermInfoGetBestCanvasMode(termInfo)
	pixelMode := chafa.TermInfoGetBestPixelMode(termInfo)

	passthrough := chafa.CHAFA_PASSTHROUGH_NONE
	if chafa.TermInfoGetIsPixelPassthroughNeeded(termInfo, pixelMode) {
		passthrough = chafa.TermInfoGetPassthroughType(termInfo)
	}

	symbolMap := chafa.SymbolMapNew()
	chafa.SymbolMapAddByTags(symbolMap, chafa.TermInfoGetSafeSymbolTags(termInfo))

	return termInfo, mode, pixelMode, passthrough, symbolMap
}

func getTTYSize() TermSize {
	termSize := TermSize{
		widthCells:   -1,
		heightCells:  -1,
		widthPixels:  -1,
		heightPixels: -1,
	}

	fds := []uintptr{os.Stdout.Fd(), os.Stderr.Fd(), os.Stdin.Fd()}
	for _, fd := range fds {
		winsize, err := unix.IoctlGetWinsize(int(fd), unix.TIOCGWINSZ)
		if err == nil {
			termSize.widthCells = int32(winsize.Col)
			termSize.heightCells = int32(winsize.Row)
			termSize.widthPixels = int32(winsize.Xpixel)
			termSize.heightPixels = int32(winsize.Ypixel)
			break
		}
	}

	if termSize.widthCells <= 0 {
		termSize.widthCells = -1
	}
	if termSize.heightCells <= 2 {
		termSize.heightCells = -1
	}

	// If .xpixel and .ypixel are filled out, we can calculate
	// aspect information for the font used. Sixel-capable terminals
	// like mlterm set these fields, but most others do not.

	if termSize.widthPixels <= 0 || termSize.heightPixels <= 0 {
		termSize.widthPixels = -1
		termSize.heightPixels = -1
	}

	return termSize
}

func convertImage(
	pixels []uint8,
	pixWidth, pixHeight, pixRowstride int32,
	pixelType chafa.PixelType,
	widthCells, heightCells, cellWidth, cellHeight int32,
) string {
	termInfo, mode, pixelMode, passthrough, symbolMap := detectTerminal()
	defer chafa.SymbolMapUnref(symbolMap)
	defer chafa.TermInfoUnref(termInfo)

	// Set up a configuration based on detected characteristics
	config := chafa.CanvasConfigNew()
	defer chafa.CanvasConfigUnref(config)

	chafa.CanvasConfigSetCanvasMode(config, mode)
	chafa.CanvasConfigSetPixelMode(config, pixelMode)
	chafa.CanvasConfigSetGeometry(config, widthCells, heightCells)
	chafa.CanvasConfigSetPassthrough(config, passthrough)
	chafa.CanvasConfigSetSymbolMap(config, symbolMap)

	if cellWidth > 0 && cellHeight > 0 {
		// We know the pixel dimensions of each cell. Store it in the config.
		chafa.CanvasConfigSetCellGeometry(config, cellWidth, cellHeight)
	}

	// Create canvas
	canvas := chafa.CanvasNew(config)
	defer chafa.CanvasUnRef(canvas)

	// Draw pixels to the canvas
	chafa.CanvasDrawAllPixels(canvas, pixelType, pixels, pixWidth, pixHeight, pixRowstride)

	// Build printable string
	printable := chafa.CanvasPrint(canvas, termInfo)

	return printable.String()
}

func main() {
	pixels := [PIX_WIDTH * PIX_HEIGHT * N_CHANNELS]uint8{
		0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
	}

	var fontRatio float32 = 0.5
	var cellWidth int32 = -1
	var cellHeight int32 = -1

	termSize := getTTYSize()

	if termSize.widthCells > 0 && termSize.heightCells > 0 &&
		termSize.widthPixels > 0 && termSize.heightPixels > 0 {
		cellWidth = termSize.widthPixels / termSize.widthCells
		cellHeight = termSize.heightPixels / termSize.heightCells
		fontRatio = float32(cellWidth) / float32(cellHeight)
	}

	widthCells := termSize.widthCells
	heightCells := termSize.heightCells

	chafa.CalcCanvasGeometry(
		PIX_HEIGHT,
		PIX_HEIGHT,
		&widthCells,
		&heightCells,
		fontRatio,
		true,
		false,
	)

	printable := convertImage(
		pixels[:],
		PIX_WIDTH,
		PIX_HEIGHT,
		PIX_WIDTH*N_CHANNELS,
		chafa.CHAFA_PIXEL_RGBA8_UNASSOCIATED,
		widthCells,
		heightCells,
		cellWidth,
		cellHeight,
	)

	fmt.Println(printable)
}
