//go:build cgo
// +build cgo

// CGo-bound helper for faults.go. The full set of *Manager PoSt methods
// live in manager_post.go itself (also //go:build cgo).
//
// Mirrors the worker_local_cgo.go pattern.

package sealer

import (
	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/storage/sealer/ffiwrapper"
	"github.com/filecoin-project/lotus/storage/sealer/storiface"
)

// generatePoStFallbackSectorChallenges wraps the ffi-backed call used
// by CheckProvable (faults.go) and manager_post.go's PoSt generation.
func generatePoStFallbackSectorChallenges(pp abi.RegisteredPoStProof, minerID abi.ActorID, randomness abi.PoStRandomness, sectorNumbers []abi.SectorNumber) (*storiface.FallbackChallenges, error) {
	return ffi.GeneratePoStFallbackSectorChallenges(pp, minerID, randomness, sectorNumbers)
}

// newProver constructs a PoSt prover backed by ffiwrapper. Used by
// Manager.New to wire the prover side of the sealer.
func newProver(rp *readonlyProvider) (storiface.ProverPoSt, error) {
	return ffiwrapper.New(rp)
}
