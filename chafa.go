package chafa

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/ebitengine/purego"
)

const libName = "libchafa"

func loadLibrary() (uintptr, error) {
	// Try to extract embedded library first
	libPath, err := extractEmbeddedLibrary()
	if err == nil {
		// Successfully extracted embedded library, try to load it
		slib, err := purego.Dlopen(libPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err == nil {
			return slib, nil
		}
		// If loading failed, log the error and fall back to system paths
		fmt.Printf("Warning: Failed to load embedded library: %v\n", err)
	} else {
		fmt.Printf("Warning: Failed to extract embedded library: %v\n", err)
	}

	// Fall back to original behavior for compatibility
	var libraryName string
	switch runtime.GOOS {
	case "darwin":
		libraryName = fmt.Sprintf("%s.dylib", libName)
	case "linux":
		libraryName = fmt.Sprintf("%s.so", libName)
	default:
		return 0, fmt.Errorf("GOOS=%s is not supported", runtime.GOOS)
	}

	libPath = os.Getenv("LD_LIBRARY_PATH")
	paths := strings.Split(libPath, ":")
	cwd, err := os.Getwd()
	if err != nil {
		return 0, err
	}
	paths = append(paths, cwd)

	for _, path := range paths {
		libPath := filepath.Join(path, libraryName)
		if _, err := os.Stat(libPath); err == nil {
			slib, dlerr := purego.Dlopen(libPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
			if dlerr != nil {
				return 0, fmt.Errorf("failed to load library at %s: %w", libPath, dlerr)
			}
			return slib, nil
		}
	}
	return 0, fmt.Errorf("%s library not found in LD_LIBRARY_PATH or CWD", libName)
}

var libOnce sync.Once

func init() {
	libOnce.Do(func() {
		libchafa, err := loadLibrary()
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

		// Config
		purego.RegisterLibFunc(&CanvasConfigNew, libchafa, "chafa_canvas_config_new")
		purego.RegisterLibFunc(&CanvasConfigCopy, libchafa, "chafa_canvas_config_copy")
		purego.RegisterLibFunc(&CanvasConfigRef, libchafa, "chafa_canvas_config_ref")
		purego.RegisterLibFunc(&CanvasConfigUnref, libchafa, "chafa_canvas_config_unref")
		purego.RegisterLibFunc(
			&CanvasConfigGetGeometry,
			libchafa,
			"chafa_canvas_config_get_geometry",
		)
		purego.RegisterLibFunc(
			&CanvasConfigSetGeometry,
			libchafa,
			"chafa_canvas_config_set_geometry",
		)
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

		// SymbolMap
		purego.RegisterLibFunc(&SymbolMapNew, libchafa, "chafa_symbol_map_new")
		purego.RegisterLibFunc(&SymbolMapCopy, libchafa, "chafa_symbol_map_copy")
		purego.RegisterLibFunc(&SymbolMapRef, libchafa, "chafa_symbol_map_ref")
		purego.RegisterLibFunc(&SymbolMapUnref, libchafa, "chafa_symbol_map_unref")
		purego.RegisterLibFunc(&SymbolMapAddByTags, libchafa, "chafa_symbol_map_add_by_tags")
		purego.RegisterLibFunc(&SymbolMapAddByRange, libchafa, "chafa_symbol_map_add_by_range")
		purego.RegisterLibFunc(&SymbolMapRemoveByTags, libchafa, "chafa_symbol_map_remove_by_tags")
		purego.RegisterLibFunc(
			&SymbolMapRemoveByRange,
			libchafa,
			"chafa_symbol_map_remove_by_range",
		)
		purego.RegisterLibFunc(
			&SymbolMapApplySelectors,
			libchafa,
			"chafa_symbol_map_apply_selectors",
		)
		purego.RegisterLibFunc(
			&SymbolMapGetAllowBuiltinGlyphs,
			libchafa,
			"chafa_symbol_map_get_allow_builtin_glyphs",
		)
		purego.RegisterLibFunc(
			&SymbolMapSetAllowBuiltinGlyphs,
			libchafa,
			"chafa_symbol_map_set_allow_builtin_glyphs",
		)
		purego.RegisterLibFunc(&SymbolMapGetGlyph, libchafa, "chafa_symbol_map_get_glyph")
		purego.RegisterLibFunc(&SymbolMapAddGlyph, libchafa, "chafa_symbol_map_add_glyph")

		// TermDb
		purego.RegisterLibFunc(&TermDbNew, libchafa, "chafa_term_db_new")
		purego.RegisterLibFunc(&TermDbCopy, libchafa, "chafa_term_db_copy")
		purego.RegisterLibFunc(&TermDbRef, libchafa, "chafa_term_db_ref")
		purego.RegisterLibFunc(&TermDbUnref, libchafa, "chafa_term_db_unref")
		purego.RegisterLibFunc(&TermDbGetDefault, libchafa, "chafa_term_db_get_default")
		purego.RegisterLibFunc(&TermDbDetect, libchafa, "chafa_term_db_detect")
		purego.RegisterLibFunc(&TermDbGetFallbackInfo, libchafa, "chafa_term_db_get_fallback_info")

		// Features
		purego.RegisterLibFunc(&GetNThreads, libchafa, "chafa_get_n_threads")
		purego.RegisterLibFunc(&GetNActualThreads, libchafa, "chafa_get_n_actual_threads")

		// Miscellaneous
		purego.RegisterLibFunc(&CalcCanvasGeometry, libchafa, "chafa_calc_canvas_geometry")
	})
}

type GString struct {
	str string
}

func (gstr *GString) String() string {
	return gstr.str
}

type GError struct {
	Domain  uint32
	Code    int32
	Message string
}

func Load(path string) (pixels []uint8, width, height int32, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	var img image.Image

	switch filepath.Ext(path) {
	case "png":
		img, err = png.Decode(file)
		if err != nil {
			return nil, 0, 0, err
		}
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
		if err != nil {
			return nil, 0, 0, err
		}
	case "gif":
		img, err = gif.Decode(file)
		if err != nil {
			return nil, 0, 0, err
		}
	default:
		img, _, err = image.Decode(file)
		if err != nil {
			return nil, 0, 0, err
		}
	}

	bounds := img.Bounds()
	width = int32(bounds.Dx())
	height = int32(bounds.Dy())

	rgbaImg := image.NewRGBA(bounds)
	draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)

	return rgbaImg.Pix, width, height, nil
}
