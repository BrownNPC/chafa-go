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

	purego.RegisterLibFunc(&CanvasNew, libchafa, "chafa_canvas_new")

	purego.RegisterLibFunc(&GetNThreads, libchafa, "chafa_get_n_threads")
	purego.RegisterLibFunc(&GetNActualThreads, libchafa, "chafa_get_n_actual_threads")
}
