//go:build !cgo
// +build !cgo

// Non-CGo stub for ffiwrapper.
//
// Under CGO_ENABLED=0 the real ffiwrapper.New (and its FFIWrapperOpts /
// FFIWrapperOpt types) live in sealer_cgo.go behind //go:build cgo.
// This file provides minimal type aliases + stub constructors so the
// surrounding sealer package compiles for PDP-only consumers (Curio
// Core). Any caller that reaches the runtime path returns a clear
// 'requires CGO_ENABLED=1' error.
//
// Mirrors the storage/paths/local_cgo.go + local_nocgo.go split.

package ffiwrapper

import (
	"errors"
)

var errFFIWrapperNotBuiltWithCGo = errors.New("ffiwrapper: this package requires building with CGO_ENABLED=1 (filecoin-ffi linkage). PDP-only deployments should never reach this code path.")

// FFIWrapperOpts is the !cgo stub of the real options struct.
// Present so external callers can write `ffiwrapper.FFIWrapperOpt`
// type references without breaking compilation.
type FFIWrapperOpts struct{}

// FFIWrapperOpt is the !cgo stub of the options-functor type. Same
// shape as the CGo version: `func(*FFIWrapperOpts)`.
type FFIWrapperOpt func(*FFIWrapperOpts)

// New under !cgo returns a stub error. Real callers (sealing workers)
// build with CGO_ENABLED=1.
func New(_ SectorProvider, _ ...FFIWrapperOpt) (*Sealer, error) {
	return nil, errFFIWrapperNotBuiltWithCGo
}
