package chafa

import (
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
)

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "/usr/lib/libchafa.dylib"
	case "linux":
		return "libchafa.so"
	default:
		panic(fmt.Errorf("GOOS=%s is not supported", runtime.GOOS))
	}
}

func init() {
	libchafa, err := purego.Dlopen(getSystemLibrary(), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// Canvas
	purego.RegisterLibFunc(&CanvasNew, libchafa, "chafa_canvas_new")
	purego.RegisterLibFunc(&CanvasNewSimilar, libchafa, "chafa_canvas_new_similar")
	purego.RegisterLibFunc(&CanvasRef, libchafa, "chafa_canvas_ref")
	purego.RegisterLibFunc(&CanvasUnRef, libchafa, "chafa_canvas_unref")
	purego.RegisterLibFunc(&CanvasPeekConfig, libchafa, "chafa_canvas_peek_config")
	purego.RegisterLibFunc(&CanvasSetPlacement, libchafa, "chafa_canvas_set_placement")
	purego.RegisterLibFunc(&CanvasDrawAllPixels, libchafa, "chafa_canvas_draw_all_pixels")
	purego.RegisterLibFunc(&CanvasPrint, libchafa, "chafa_canvas_print")
	purego.RegisterLibFunc(&CanvasPrintRows, libchafa, "chafa_canvas_print_rows")
	purego.RegisterLibFunc(&CanvasPrintRowsStrv, libchafa, "chafa_canvas_print_rows_strv")
	purego.RegisterLibFunc(&CanvasGetCharAt, libchafa, "chafa_canvas_get_char_at")
	purego.RegisterLibFunc(&CanvasSetCharAt, libchafa, "chafa_canvas_set_char_at")
	purego.RegisterLibFunc(&CanvasGetColorsAt, libchafa, "chafa_canvas_get_colors_at")
	purego.RegisterLibFunc(&CanvasSetColorsAt, libchafa, "chafa_canvas_set_colors_at")
	purego.RegisterLibFunc(&CanvasGetRawColorsAt, libchafa, "chafa_canvas_get_raw_colors_at")
	purego.RegisterLibFunc(&CanvasSetRawColorsAt, libchafa, "chafa_canvas_set_raw_colors_at")

	// SymbolMap
	purego.RegisterLibFunc(&SymbolMapNew, libchafa, "chafa_symbol_map_new")
	purego.RegisterLibFunc(&SymbolMapUnref, libchafa, "chafa_symbol_map_unref")
	purego.RegisterLibFunc(&SymbolMapAddByTags, libchafa, "chafa_symbol_map_add_by_tags")

	// Config
	purego.RegisterLibFunc(&CanvasConfigNew, libchafa, "chafa_canvas_config_new")
	purego.RegisterLibFunc(&CanvasConfigCopy, libchafa, "chafa_canvas_config_copy")
	purego.RegisterLibFunc(&CanvasConfigRef, libchafa, "chafa_canvas_config_ref")
	purego.RegisterLibFunc(&CanvasConfigUnref, libchafa, "chafa_canvas_config_unref")
	purego.RegisterLibFunc(&CanvasConfigGetGeometry, libchafa, "chafa_canvas_config_get_geometry")
	purego.RegisterLibFunc(&CanvasConfigSetGeometry, libchafa, "chafa_canvas_config_set_geometry")
	purego.RegisterLibFunc(
		&CanvasConfigGetCellGeometry,
		libchafa,
		"chafa_canvas_config_get_cell_geometry",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetCellGeometry,
		libchafa,
		"chafa_canvas_config_set_cell_geometry",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetPixelMode,
		libchafa,
		"chafa_canvas_config_get_pixel_mode",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetPixelMode,
		libchafa,
		"chafa_canvas_config_set_pixel_mode",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetCanvasMode,
		libchafa,
		"chafa_canvas_config_get_canvas_mode",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetCanvasMode,
		libchafa,
		"chafa_canvas_config_set_canvas_mode",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetColorExtractor,
		libchafa,
		"chafa_canvas_config_get_color_extractor",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetColorExtractor,
		libchafa,
		"chafa_canvas_config_set_color_extractor",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetColorSpace,
		libchafa,
		"chafa_canvas_config_get_color_space",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetColorSpace,
		libchafa,
		"chafa_canvas_config_set_color_space",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetPreprocessingEnabled,
		libchafa,
		"chafa_canvas_config_get_preprocessing_enabled",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetPreprocessingEnabled,
		libchafa,
		"chafa_canvas_config_set_preprocessing_enabled",
	)
	purego.RegisterLibFunc(
		&CanvasConfigPeekSymbolMap,
		libchafa,
		"chafa_canvas_config_peek_symbol_map",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetSymbolMap,
		libchafa,
		"chafa_canvas_config_set_symbol_map",
	)
	purego.RegisterLibFunc(
		&CanvasConfigPeekFillSymbolMap,
		libchafa,
		"chafa_canvas_config_peek_fill_symbol_map",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetTransparencyThreshold,
		libchafa,
		"chafa_canvas_config_get_transparency_threshold",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetTransparencyThreshold,
		libchafa,
		"chafa_canvas_config_set_transparency_threshold",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetFgOnlyEnabled,
		libchafa,
		"chafa_canvas_config_get_fg_only_enabled",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetFgOnlyEnabled,
		libchafa,
		"chafa_canvas_config_set_fg_only_enabled",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetFgColor,
		libchafa,
		"chafa_canvas_config_get_fg_color",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetFgColor,
		libchafa,
		"chafa_canvas_config_set_fg_color",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetBgColor,
		libchafa,
		"chafa_canvas_config_get_bg_color",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetBgColor,
		libchafa,
		"chafa_canvas_config_set_bg_color",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetWorkFactor,
		libchafa,
		"chafa_canvas_config_get_work_factor",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetWorkFactor,
		libchafa,
		"chafa_canvas_config_set_work_factor",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetDitherMode,
		libchafa,
		"chafa_canvas_config_get_dither_mode",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetDitherMode,
		libchafa,
		"chafa_canvas_config_set_dither_mode",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetDitherGrainSize,
		libchafa,
		"chafa_canvas_config_get_dither_grain_size",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetDitherGrainSize,
		libchafa,
		"chafa_canvas_config_set_dither_grain_size",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetDitherIntensity,
		libchafa,
		"chafa_canvas_config_get_dither_intensity",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetDitherIntensity,
		libchafa,
		"chafa_canvas_config_set_dither_intensity",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetOptimizations,
		libchafa,
		"chafa_canvas_config_get_optimizations",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetOptimizations,
		libchafa,
		"chafa_canvas_config_set_optimizations",
	)
	purego.RegisterLibFunc(
		&CanvasConfigGetPassthrough,
		libchafa,
		"chafa_canvas_config_get_passthrough",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetPassthrough,
		libchafa,
		"chafa_canvas_config_set_passthrough",
	)

	// Placement
	purego.RegisterLibFunc(&PlacementNew, libchafa, "chafa_placement_new")
	purego.RegisterLibFunc(&PlacementRef, libchafa, "chafa_placement_ref")
	purego.RegisterLibFunc(&PlacementUnref, libchafa, "chafa_placement_unref")
	purego.RegisterLibFunc(&PlacementGetTuck, libchafa, "chafa_placement_get_tuck")
	purego.RegisterLibFunc(&PlacementSetTuck, libchafa, "chafa_placement_set_tuck")
	purego.RegisterLibFunc(&PlacementGetHAlign, libchafa, "chafa_placement_get_halign")
	purego.RegisterLibFunc(&PlacementSetHAlign, libchafa, "chafa_placement_set_halign")
	purego.RegisterLibFunc(&PlacementGetVAlign, libchafa, "chafa_placement_get_valign")
	purego.RegisterLibFunc(&PlacementSetVAlign, libchafa, "chafa_placement_set_valign")

	// Image
	purego.RegisterLibFunc(&ImageNew, libchafa, "chafa_image_new")
	purego.RegisterLibFunc(&ImageRef, libchafa, "chafa_image_ref")
	purego.RegisterLibFunc(&ImageUnref, libchafa, "chafa_image_unref")
	purego.RegisterLibFunc(&ImageSetFrame, libchafa, "chafa_image_set_frame")

	// Frame
	purego.RegisterLibFunc(&FrameNew, libchafa, "chafa_frame_new")
	purego.RegisterLibFunc(&FrameNewBorrow, libchafa, "chafa_frame_new_borrow")
	purego.RegisterLibFunc(&FrameNewSteal, libchafa, "chafa_frame_new_steal")
	purego.RegisterLibFunc(&FrameRef, libchafa, "chafa_frame_ref")
	purego.RegisterLibFunc(&FrameUnref, libchafa, "chafa_frame_unref")

	// Features
	purego.RegisterLibFunc(&GetNThreads, libchafa, "chafa_get_n_threads")
	purego.RegisterLibFunc(&GetNActualThreads, libchafa, "chafa_get_n_actual_threads")
}

type GString struct {
	str string
}

func (gstr *GString) String() string {
	return gstr.str
}
