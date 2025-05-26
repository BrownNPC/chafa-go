package chafa

import "unsafe"

var (
	// Creates a new, blank [TermDb].
	TermDbNew func() *TermDb

	// Creates a new [TermDb] that's a copy of termDb.
	TermDbCopy func(termDb *TermDb) *TermDb

	// Adds a reference to termDb.
	TermDbRef func(termDb *TermDb)

	// Removes a reference from termDb.
	TermDbUnref func(termDb *TermDb)

	// Gets the global [TermDb]. This can normally be used safely in a read-only
	// capacity. The caller should not unref the returned object.
	TermDbGetDefault func() *TermDb

	// Builds a new [TermInfo] with capabilities implied by the provided
	// environment variables (principally the TERM variable, but also others).
	termDbDetect func(termDb *TermDb, envp **byte) *TermInfo

	// Builds a new [TermInfo] with fallback control sequences. This can be used
	// with unknown but presumably modern terminals, or to supplement missing
	// capabilities in a detected terminal.
	//
	// Fallback control sequences may cause unpredictable behavior and should
	// only be used as a last resort.
	TermDbGetFallbackInfo func(termDb *TermDb) *TermInfo
)

func TermDbDetect(termDb *TermDb, envp []string) *TermInfo {
	ptrs := make([]*byte, len(envp))
	allocated := make([][]byte, len(envp))

	for i, s := range envp {
		cstr := append([]byte(s), 0)
		allocated[i] = cstr
		ptrs[i] = &cstr[0]
	}

	ptrBlock := make([]uintptr, len(ptrs))
	for i, p := range ptrs {
		ptrBlock[i] = uintptr(unsafe.Pointer(p))
	}

	return termDbDetect(termDb, (**byte)(unsafe.Pointer(&ptrBlock[0])))
}

type TermDb struct {
	Refs int32
}
