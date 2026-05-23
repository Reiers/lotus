//go:build cgo
// +build cgo

// CGo-bound helpers for worker_local.go.
//
// FFIExec wires ffiwrapper.New (which calls into filecoin-ffi) into the
// Storage interface. localWorkerGPUDevices probes the host for GPU
// devices via ffi.GetGPUDevices.
//
// These functions live in their own file under //go:build cgo so that
// storage/sealer can compile under CGO_ENABLED=0 with stubs in
// worker_local_nocgo.go. PDP-only consumers (Curio Core) never reach
// these code paths; sealing workers do.

package sealer

import (
	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/storage/sealer/ffiwrapper"
	"github.com/filecoin-project/lotus/storage/sealer/storiface"
)

// FFIExec returns a LocalWorker -> Storage constructor backed by
// ffiwrapper. CGo-only.
func FFIExec(opts ...ffiwrapper.FFIWrapperOpt) func(l *LocalWorker) (storiface.Storage, error) {
	return func(l *LocalWorker) (storiface.Storage, error) {
		return ffiwrapper.New(&localWorkerPathProvider{w: l}, opts...)
	}
}

// localWorkerGPUDevices probes the host for GPU devices via
// filecoin-ffi. Returns an empty slice on probe error (Worker.Info
// continues with a warning log via the upstream call site).
func localWorkerGPUDevices() []string {
	gpus, err := ffi.GetGPUDevices()
	if err != nil {
		log.Errorf("getting gpu devices failed: %+v", err)
		return nil
	}
	return gpus
}
