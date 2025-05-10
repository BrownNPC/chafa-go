package chafa

var (
	CanvasNew func(config *CanvasConfig) *Canvas
)

type Canvas struct {
	Refs int32

	WidthPixels, HeightPixels int32
	// ChafaPixel *pixels;
	// ChafaCanvasCell *cells;
	// guint have_alpha : 1;
	// guint needs_clear : 1;
	//
	// /* Whether to consider inverted symbols; FALSE if using FG only */
	// guint consider_inverted : 1;
	//
	// /* Whether to extract symbol colors; FALSE if using default colors */
	// guint extract_colors : 1;
	//
	// /* Whether to quantize colors before calculating error (slower, but
	//  * yields better results in palettized modes, especially 16/8) */
	// guint use_quantized_error : 1;
	//
	// ChafaColorPair default_colors;
	// guint work_factor_int;
	//
	// /* Character to use in cells where fg color == bg color. Typically
	//  * space, but could be something else depending on the symbol map. */
	// gunichar blank_char;
	//
	// /* Character to use in cells where fg color == bg color and the color
	//  * is only legal in FG. Typically 0x2588 (solid block), but could be
	//  * something else depending on the symbol map. Can be zero if there is
	//  * no good candidate! */
	// gunichar solid_char;
	//
	// ChafaCanvasConfig config;
	//
	// /* Used when setting pixel data */
	// ChafaDither dither;
	//
	// /* This is NULL in CHAFA_PIXEL_MODE_SYMBOLS, otherwise one of:
	//  * (ChafaSixelCanvas *), (ChafaKittyCanvas *), (ChafaIterm2Canvas *) */
	// gpointer pixel_canvas;
	//
	// /* It's possible to have a single placement that covers the entire
	//  * canvas. In this case, it is stored here. */
	// ChafaPlacement *placement;
	//
	// /* Our palettes. Kind of a big structure, so they go last. */
	// ChafaPalette fg_palette;
	// ChafaPalette bg_palette;
}

type CanvasConfig struct {
	Refs int32

	Width, Height int32
	//   gint cell_width, cell_height;
	//   ChafaCanvasMode canvas_mode;
	//   ChafaColorSpace color_space;
	//   ChafaDitherMode dither_mode;
	//   ChafaColorExtractor color_extractor;
	//   ChafaPixelMode pixel_mode;
	//   gint dither_grain_width, dither_grain_height;
	//   gfloat dither_intensity;
	//   guint32 fg_color_packed_rgb;
	//   guint32 bg_color_packed_rgb;
	//   gint alpha_threshold;  /* 0-255. 255 = no alpha in output */
	//   gfloat work_factor;
	//   ChafaSymbolMap symbol_map;
	//   ChafaSymbolMap fill_symbol_map;
	//   guint preprocessing_enabled : 1;
	//   guint fg_only_enabled : 1;
	//   ChafaOptimizations optimizations;
	//   ChafaPassthrough passthrough;
}
