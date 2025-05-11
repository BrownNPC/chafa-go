package chafa

import "unsafe"

var (
	CanvasNew           func(config *CanvasConfig) *Canvas
	CanvasNewSimilar    func(orig *Canvas) *Canvas
	CanvasRef           func(canvas *Canvas)
	CanvasUnRef         func(canvas *Canvas)
	CanvasPeekConfig    func(canvas *Canvas) *CanvasConfig
	CanvasSetPlacement  func(canvas *Canvas, placement *Placement)
	CanvasDrawAllPixels func(canvas *Canvas,
		src_pixel_type PixelType,
		src_pixels []uint8,
		src_width int32,
		src_height int32,
		src_rowstride int32,
	)
	CanvasPrint     func(canvas *Canvas, term_info *TermInfo) *GString
	CanvasPrintRows func(canvas *Canvas, term_info *TermInfo, array_out *[]string, array_len_out *int32) // TODO:
)

type Canvas struct {
	Refs int32

	WidthPixels, HeightPixels int32
	Pixels                    *Pixel
	Cells                     *CanvasCell
	HaveAlpha                 bool
	NeedsClear                bool

	// Whether to consider inverted symbols; FALSE if using FG only
	ConsiderInverted bool

	// Whether to extract symbol colors; FALSE if using default colors
	ExtractColors bool

	// Whether to quantize colors before calculating error (slower, but
	// yields better results in palettized modes, especially 16/8)
	UseQuantizedError bool

	DefaultColors ColorPair
	WorkFactorInt uint32

	// Character to use in cells where fg color == bg color. Typically
	// space, but could be something else depending on the symbol map.
	BlankChar uint8 // gunichar

	// Character to use in cells where fg color == bg color and the color
	// is only legal in FG. Typically 0x2588 (solid block), but could be
	// something else depending on the symbol map. Can be zero if there is
	// no good candidate!
	SolidChar uint8 // gunichar

	Config CanvasConfig

	// Used when setting pixel data
	Dither Dither

	// This is NULL in CHAFA_PIXEL_MODE_SYMBOLS, otherwise one of:
	// (ChafaSixelCanvas *), (ChafaKittyCanvas *), (ChafaIterm2Canvas *)
	PixelCanvas unsafe.Pointer

	// It's possible to have a single placement that covers the entire
	// canvas. In this case, it is stored here.
	Placement *Placement

	// Our palettes. Kind of a big structure, so they go last.
	FgPalette Palette
	BgPalette Palette
}

type Pixel struct {
	Col Color
}

type Color struct {
	Ch [4]uint8
}

type CanvasCell struct {
	C uint8

	// Colors can be either packed RGBA or index
	FgColor uint32
	BgColor uint32
}

type ColorPair struct {
	Colors [2]Color
}

type Dither struct {
	Mode             DitherMode
	Intensity        float64
	GrainWidthShift  int32
	GrainHeightShift int32

	TextureSizeShift int32
	TextureSizeMask  uint32
	TextureData      *int32
}

type DitherMode int32

const (
	CHAFA_DITHER_MODE_NONE      DitherMode = 0
	CHAFA_DITHER_MODE_ORDERED   DitherMode = 1
	CHAFA_DITHER_MODE_DIFFUSION DitherMode = 2
	CHAFA_DITHER_MODE_NOISE     DitherMode = 3
	CHAFA_DITHER_MODE_MAX       DitherMode = 4
)

type Placement struct {
	Refs int32

	Image          *Image
	Id             int32
	Halign, Valign Align
	Tuck           Tuck
}

type Image struct {
	Refs  int32
	Frame *Frame
}

type Frame struct {
	Refs                     int32
	PixelType                PixelType
	Width, Height, Rowstride int32

	Data unsafe.Pointer

	DataIsOwned bool
}

type PixelType int32

const (
	/* 32 bits per pixel */

	CHAFA_PIXEL_RGBA8_PREMULTIPLIED PixelType = 0
	CHAFA_PIXEL_BGRA8_PREMULTIPLIED PixelType = 1
	CHAFA_PIXEL_ARGB8_PREMULTIPLIED PixelType = 2
	CHAFA_PIXEL_ABGR8_PREMULTIPLIED PixelType = 3

	CHAFA_PIXEL_RGBA8_UNASSOCIATED PixelType = 4
	CHAFA_PIXEL_BGRA8_UNASSOCIATED PixelType = 5
	CHAFA_PIXEL_ARGB8_UNASSOCIATED PixelType = 6
	CHAFA_PIXEL_ABGR8_UNASSOCIATED PixelType = 7

	/* 24 bits per pixel */

	CHAFA_PIXEL_RGB8 PixelType = 8
	CHAFA_PIXEL_BGR8 PixelType = 9

	CHAFA_PIXEL_MAX PixelType = 10
)

type Align int32

const (
	CHAFA_ALIGN_START  Align = 0
	CHAFA_ALIGN_END    Align = 1
	CHAFA_ALIGN_CENTER Align = 2
	CHAFA_ALIGN_MAX    Align = 3
)

type Tuck int32

const (
	CHAFA_TUCK_STRETCH       Tuck = 0
	CHAFA_TUCK_FIT           Tuck = 1
	CHAFA_TUCK_SHRINK_TO_FIT Tuck = 2
	CHAFA_TUCK_MAX           Tuck = 3
)

const CHAFA_PALETTE_INDEX_MAX = 259

type Palette struct {
	Type             PaletteType
	Colors           [CHAFA_PALETTE_INDEX_MAX]PaletteColor
	Table            [CHAFA_COLOR_SPACE_MAX]ColorTable
	FirstColor       int32
	NColors          int32
	AlphaThreshold   int32
	TransparentIndex int32
}

type PaletteType int32

const (
	CHAFA_PALETTE_TYPE_DYNAMIC_256 PaletteType = 0
	CHAFA_PALETTE_TYPE_FIXED_256   PaletteType = 1
	CHAFA_PALETTE_TYPE_FIXED_240   PaletteType = 2
	CHAFA_PALETTE_TYPE_FIXED_16    PaletteType = 3
	CHAFA_PALETTE_TYPE_FIXED_8     PaletteType = 4
	CHAFA_PALETTE_TYPE_FIXED_FGBG  PaletteType = 5
)

type PaletteColor struct {
	Col [CHAFA_COLOR_SPACE_MAX]Color
}

type ColorSpace int32

const (
	CHAFA_COLOR_SPACE_RGB    ColorSpace = 0
	CHAFA_COLOR_SPACE_DIN99D ColorSpace = 1
	CHAFA_COLOR_SPACE_MAX    ColorSpace = 2
)

const CHAFA_COLOR_TABLE_MAX_ENTRIES = 256

type ColorTableEntry struct {
	V   [2]int32
	Pen int32
}

type ColorTable struct {
	Entries [CHAFA_COLOR_TABLE_MAX_ENTRIES]ColorTableEntry

	// Each pen is 24 bits (B8G8R8) of color information
	Pens [CHAFA_COLOR_TABLE_MAX_ENTRIES]uint32

	NEntries int32
	IsSorted bool

	Eigenvectors [2]Vec3i32
	Average      Vec3i32

	EigenMul [2]uint32
}

type Vec3i32 struct {
	V [3]int32
}

type CanvasMode int32

const (
	CHAFA_CANVAS_MODE_TRUECOLOR    CanvasMode = 0
	CHAFA_CANVAS_MODE_INDEXED_256  CanvasMode = 1
	CHAFA_CANVAS_MODE_INDEXED_240  CanvasMode = 2
	CHAFA_CANVAS_MODE_INDEXED_16   CanvasMode = 3
	CHAFA_CANVAS_MODE_FGBG_BGFG    CanvasMode = 4
	CHAFA_CANVAS_MODE_FGBG         CanvasMode = 5
	CHAFA_CANVAS_MODE_INDEXED_8    CanvasMode = 6
	CHAFA_CANVAS_MODE_INDEXED_16_8 CanvasMode = 7
	CHAFA_CANVAS_MODE_MAX          CanvasMode = 8
)

type ColorExtractor int32

const (
	CHAFA_COLOR_EXTRACTOR_AVERAGE ColorExtractor = 0
	CHAFA_COLOR_EXTRACTOR_MEDIAN  ColorExtractor = 1
	CHAFA_COLOR_EXTRACTOR_MAX     ColorExtractor = 2
)

type PixelMode int32

const (
	CHAFA_PIXEL_MODE_SYMBOLS PixelMode = 0
	CHAFA_PIXEL_MODE_SIXELS  PixelMode = 1
	CHAFA_PIXEL_MODE_KITTY   PixelMode = 2
	CHAFA_PIXEL_MODE_ITERM2  PixelMode = 3
	CHAFA_PIXEL_MODE_MAX     PixelMode = 4
)

type Optimizations int32

const (
	CHAFA_OPTIMIZATION_REUSE_ATTRIBUTES Optimizations = (1 << 0)
	CHAFA_OPTIMIZATION_SKIP_CELLS       Optimizations = (1 << 1)
	CHAFA_OPTIMIZATION_REPEAT_CELLS     Optimizations = (1 << 2)
	CHAFA_OPTIMIZATION_NONE             Optimizations = 0
	CHAFA_OPTIMIZATION_ALL              Optimizations = 0x7fffffff
)

type Passthrough int32

const (
	CHAFA_PASSTHROUGH_NONE   Passthrough = 0
	CHAFA_PASSTHROUGH_SCREEN Passthrough = 1
	CHAFA_PASSTHROUGH_TMUX   Passthrough = 2
	CHAFA_PASSTHROUGH_MAX    Passthrough = 3
)
