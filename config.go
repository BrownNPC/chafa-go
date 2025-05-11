package chafa

var (
	CanvasConfigNew          func() *CanvasConfig
	CanvasConfigUnref        func(config *CanvasConfig)
	CanvasConfigSetGeometry  func(config *CanvasConfig, width, height int32)
	CanvasConfigSetSymbolMap func(config *CanvasConfig, symbolMap *SymbolMap)
	CanvasConfigSetPixelMode func(config *CanvasConfig, pixelMode PixelMode)
)

type CanvasConfig struct {
	Refs int32

	Width, Height                       int32
	CellWidth, CellHeight               int32
	CanvasMode                          CanvasMode
	ColorSpace                          ColorSpace
	DitherMode                          DitherMode
	ColorExtractor                      ColorExtractor
	PixelMode                           PixelMode
	DitherGrainWidth, DitherGrainHeight int32
	DitherIntensity                     float32
	FgColorPackedRgb                    uint32
	BgColorPackedRgb                    uint32
	AlphaThreshold                      int32 // 0-255. 255 = no alpha in output
	WorkFactor                          float32
	SymbolMap                           SymbolMap
	FillSymbolMap                       SymbolMap
	PreprocessingEnabled                bool
	FgOnlyEnabled                       bool
	Optimizations                       Optimizations
	Passthrough                         Passthrough
}
