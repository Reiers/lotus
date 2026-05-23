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
	"context"
	"errors"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/proof"

	"github.com/filecoin-project/lotus/storage/sealer/storiface"
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

// nocgoProver / nocgoVerifier are the !cgo stubs for ProofProver /
// ProofVerifier package-level singletons. All methods return
// errFFIWrapperNotBuiltWithCGo. Curio Core never reaches these code
// paths; sealing-side callers build with CGO_ENABLED=1.

type nocgoProver struct{}

var ProofProver = nocgoProver{}

var _ storiface.Prover = ProofProver

func (nocgoProver) AggregateSealProofs(_ proof.AggregateSealVerifyProofAndInfos, _ [][]byte) ([]byte, error) {
	return nil, errFFIWrapperNotBuiltWithCGo
}

type nocgoVerifier struct{}

var ProofVerifier = nocgoVerifier{}

var _ storiface.Verifier = ProofVerifier

func (nocgoVerifier) VerifySeal(_ proof.SealVerifyInfo) (bool, error) {
	return false, errFFIWrapperNotBuiltWithCGo
}
func (nocgoVerifier) VerifyAggregateSeals(_ proof.AggregateSealVerifyProofAndInfos) (bool, error) {
	return false, errFFIWrapperNotBuiltWithCGo
}
func (nocgoVerifier) VerifyReplicaUpdate(_ proof.ReplicaUpdateInfo) (bool, error) {
	return false, errFFIWrapperNotBuiltWithCGo
}
func (nocgoVerifier) VerifyWinningPoSt(_ context.Context, _ proof.WinningPoStVerifyInfo) (bool, error) {
	return false, errFFIWrapperNotBuiltWithCGo
}
func (nocgoVerifier) VerifyWindowPoSt(_ context.Context, _ proof.WindowPoStVerifyInfo) (bool, error) {
	return false, errFFIWrapperNotBuiltWithCGo
}
func (nocgoVerifier) GenerateWinningPoStSectorChallenge(_ context.Context, _ abi.RegisteredPoStProof, _ abi.ActorID, _ abi.PoStRandomness, _ uint64) ([]uint64, error) {
	return nil, errFFIWrapperNotBuiltWithCGo
}
