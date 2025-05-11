package chafa

import "unsafe"

var (
	SymbolMapNew       func() *SymbolMap
	SymbolMapAddByTags func(symbolMap *SymbolMap, tags SymbolTags)
	SymbolMapUnref     func(symbol_map *SymbolMap)
)

type SymbolMap struct {
	Refs int32

	NeedRebuild      bool
	UseBuiltinGlyphs bool

	Glyphs    unsafe.Pointer
	Glyphs2   unsafe.Pointer // Wide glyphs with left/right bitmaps
	Selectors unsafe.Pointer

	// /* Remaining fields are populated by chafa_symbol_map_prepare () */

	// Narrow symbols
	Symbols       []Symbol
	NSymbols      int32
	PackedBitmaps *uint64

	// Wide symbols
	Symbols2       []Symbol2
	NSymbols2      int32
	PackedBitmaps2 *uint64
}

type Symbol struct {
	Sc                 SymbolTags
	C                  uint8
	Coverage           byte //gchar *
	MaskU32            *uint32
	FgWeight, BgWeight int32
	Bitmap             uint64
	Popcount           int32
}

type Symbol2 struct {
	Sym [2]Symbol
}

type SymbolTags int32

const (
	CHAFA_SYMBOL_TAG_NONE      SymbolTags = 0
	CHAFA_SYMBOL_TAG_SPACE     SymbolTags = (1 << 0)
	CHAFA_SYMBOL_TAG_SOLID     SymbolTags = (1 << 1)
	CHAFA_SYMBOL_TAG_STIPPLE   SymbolTags = (1 << 2)
	CHAFA_SYMBOL_TAG_BLOCK     SymbolTags = (1 << 3)
	CHAFA_SYMBOL_TAG_BORDER    SymbolTags = (1 << 4)
	CHAFA_SYMBOL_TAG_DIAGONAL  SymbolTags = (1 << 5)
	CHAFA_SYMBOL_TAG_DOT       SymbolTags = (1 << 6)
	CHAFA_SYMBOL_TAG_QUAD      SymbolTags = (1 << 7)
	CHAFA_SYMBOL_TAG_HHALF     SymbolTags = (1 << 8)
	CHAFA_SYMBOL_TAG_VHALF     SymbolTags = (1 << 9)
	CHAFA_SYMBOL_TAG_HALF      SymbolTags = ((CHAFA_SYMBOL_TAG_HHALF) | (CHAFA_SYMBOL_TAG_VHALF))
	CHAFA_SYMBOL_TAG_INVERTED  SymbolTags = (1 << 10)
	CHAFA_SYMBOL_TAG_BRAILLE   SymbolTags = (1 << 11)
	CHAFA_SYMBOL_TAG_TECHNICAL SymbolTags = (1 << 12)
	CHAFA_SYMBOL_TAG_GEOMETRIC SymbolTags = (1 << 13)
	CHAFA_SYMBOL_TAG_ASCII     SymbolTags = (1 << 14)
	CHAFA_SYMBOL_TAG_ALPHA     SymbolTags = (1 << 15)
	CHAFA_SYMBOL_TAG_DIGIT     SymbolTags = (1 << 16)
	CHAFA_SYMBOL_TAG_ALNUM     SymbolTags = CHAFA_SYMBOL_TAG_ALPHA | CHAFA_SYMBOL_TAG_DIGIT
	CHAFA_SYMBOL_TAG_NARROW    SymbolTags = (1 << 17)
	CHAFA_SYMBOL_TAG_WIDE      SymbolTags = (1 << 18)
	CHAFA_SYMBOL_TAG_AMBIGUOUS SymbolTags = (1 << 19)
	CHAFA_SYMBOL_TAG_UGLY      SymbolTags = (1 << 20)
	CHAFA_SYMBOL_TAG_LEGACY    SymbolTags = (1 << 21)
	CHAFA_SYMBOL_TAG_SEXTANT   SymbolTags = (1 << 22)
	CHAFA_SYMBOL_TAG_WEDGE     SymbolTags = (1 << 23)
	CHAFA_SYMBOL_TAG_LATIN     SymbolTags = (1 << 24)
	CHAFA_SYMBOL_TAG_IMPORTED  SymbolTags = (1 << 25)
	CHAFA_SYMBOL_TAG_OCTANT    SymbolTags = (1 << 26)
	CHAFA_SYMBOL_TAG_EXTRA     SymbolTags = (1 << 30)
	CHAFA_SYMBOL_TAG_BAD       SymbolTags = CHAFA_SYMBOL_TAG_AMBIGUOUS | CHAFA_SYMBOL_TAG_UGLY
	CHAFA_SYMBOL_TAG_ALL       SymbolTags = ^(CHAFA_SYMBOL_TAG_EXTRA | CHAFA_SYMBOL_TAG_BAD)
)
