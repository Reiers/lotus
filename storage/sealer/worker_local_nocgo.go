//go:build !cgo
// +build !cgo

// Non-CGo stubs for worker_local.go.
//
// These keep storage/sealer compiling under CGO_ENABLED=0 so that
// downstream consumers that don't run a sealing worker (notably Curio
// Core's PDP-only deployment shape) can build a pure-Go binary. Any
// runtime caller hitting these stubs gets a clear error pointing at
// the build configuration.
//
// Mirrors the split in storage/paths/local_cgo.go + local_nocgo.go.

package sealer

import (
	"errors"

	"github.com/filecoin-project/lotus/storage/sealer/ffiwrapper"
	"github.com/filecoin-project/lotus/storage/sealer/storiface"
)

var errWorkerNotBuiltWithCGo = errors.New("sealer/worker_local: this method requires building with CGO_ENABLED=1 (filecoin-ffi linkage). PDP-only deployments should never call FFIExec.")

// FFIExec under !cgo returns a constructor that fails fast at the
// moment a LocalWorker is being wired against ffiwrapper. Curio Core's
// scheduler never calls this; full Curio with sealing does.
func FFIExec(opts ...ffiwrapper.FFIWrapperOpt) func(l *LocalWorker) (storiface.Storage, error) {
	_ = opts
	return func(l *LocalWorker) (storiface.Storage, error) {
		return nil, errWorkerNotBuiltWithCGo
	}
}

// localWorkerGPUDevices returns nil under !cgo. The caller (Worker.Info)
// reports "0 GPU devices" via its existing log line, which is the
// correct semantic for a PDP-only deployment that has no sealing GPUs.
func localWorkerGPUDevices() []string {
	return nil
}
