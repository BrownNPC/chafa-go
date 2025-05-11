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

	// SymbolMap
	purego.RegisterLibFunc(&SymbolMapNew, libchafa, "chafa_symbol_map_new")
	purego.RegisterLibFunc(&SymbolMapUnref, libchafa, "chafa_symbol_map_unref")
	purego.RegisterLibFunc(&SymbolMapAddByTags, libchafa, "chafa_symbol_map_add_by_tags")

	// Config
	purego.RegisterLibFunc(&CanvasConfigNew, libchafa, "chafa_canvas_config_new")
	purego.RegisterLibFunc(&CanvasConfigUnref, libchafa, "chafa_canvas_config_unref")
	purego.RegisterLibFunc(&CanvasConfigSetGeometry, libchafa, "chafa_canvas_config_set_geometry")
	purego.RegisterLibFunc(
		&CanvasConfigSetSymbolMap,
		libchafa,
		"chafa_canvas_config_set_symbol_map",
	)
	purego.RegisterLibFunc(
		&CanvasConfigSetPixelMode,
		libchafa,
		"chafa_canvas_config_set_pixel_mode",
	)

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
