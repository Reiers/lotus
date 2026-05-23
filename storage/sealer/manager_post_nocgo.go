//go:build !cgo
// +build !cgo

// Non-CGo stubs for manager_post.go.
//
// The PoSt generation pipeline (GenerateWinningPoSt, GenerateWindowPoSt,
// the various *WithVanilla variants) requires filecoin-ffi. Under !cgo
// the entire file is excluded and stubs here keep the package compiling
// for PDP-only consumers (Curio Core) that never run sealing PoSt.
//
// Callers that DO run sealing (lotus-miner, lotus-worker, the
// wdpost/winning_prover packages) build with CGO_ENABLED=1 and link
// the real manager_post.go.
//
// Mirrors the worker_local_cgo.go / worker_local_nocgo.go split.

package sealer

import (
	"context"
	"errors"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/proof"

	"github.com/filecoin-project/lotus/storage/sealer/storiface"
)

var errPostNotBuiltWithCGo = errors.New("sealer/manager_post: PoSt generation requires CGO_ENABLED=1 (filecoin-ffi linkage). PDP-only deployments should never call these methods.")

// generatePoStFallbackSectorChallenges under !cgo. CheckProvable still
// runs but every sector is marked unprovable (which is the truthful
// answer for a daemon that can't run PoSt anyway).
func generatePoStFallbackSectorChallenges(_ abi.RegisteredPoStProof, _ abi.ActorID, _ abi.PoStRandomness, _ []abi.SectorNumber) (*storiface.FallbackChallenges, error) {
	return nil, errPostNotBuiltWithCGo
}

// nocgoProver is a placeholder Prover that returns errPostNotBuiltWithCGo
// on every PoSt-shaped method. Allows Manager.New to construct
// successfully under !cgo even though PoSt generation can't run.
type nocgoProver struct{}

func (nocgoProver) GenerateWinningPoSt(_ context.Context, _ abi.ActorID, _ []proof.ExtendedSectorInfo, _ abi.PoStRandomness) ([]proof.PoStProof, error) {
	return nil, errPostNotBuiltWithCGo
}
func (nocgoProver) GenerateWindowPoSt(_ context.Context, _ abi.ActorID, _ abi.RegisteredPoStProof, _ []proof.ExtendedSectorInfo, _ abi.PoStRandomness) ([]proof.PoStProof, []abi.SectorID, error) {
	return nil, nil, errPostNotBuiltWithCGo
}
func (nocgoProver) GenerateWinningPoStWithVanilla(_ context.Context, _ abi.RegisteredPoStProof, _ abi.ActorID, _ abi.PoStRandomness, _ [][]byte) ([]proof.PoStProof, error) {
	return nil, errPostNotBuiltWithCGo
}
func (nocgoProver) GenerateWindowPoStWithVanilla(_ context.Context, _ abi.RegisteredPoStProof, _ abi.ActorID, _ abi.PoStRandomness, _ [][]byte, _ int) (proof.PoStProof, error) {
	return proof.PoStProof{}, errPostNotBuiltWithCGo
}

func newProver(_ *readonlyProvider) (storiface.ProverPoSt, error) {
	return nocgoProver{}, nil
}

func (m *Manager) GenerateWinningPoSt(_ context.Context, _ abi.ActorID, _ []proof.ExtendedSectorInfo, _ abi.PoStRandomness) ([]proof.PoStProof, error) {
	return nil, errPostNotBuiltWithCGo
}

func (m *Manager) GenerateWindowPoSt(_ context.Context, _ abi.ActorID, _ abi.RegisteredPoStProof, _ []proof.ExtendedSectorInfo, _ abi.PoStRandomness) ([]proof.PoStProof, []abi.SectorID, error) {
	return nil, nil, errPostNotBuiltWithCGo
}

func (m *Manager) GenerateWinningPoStWithVanilla(_ context.Context, _ abi.RegisteredPoStProof, _ abi.ActorID, _ abi.PoStRandomness, _ [][]byte) ([]proof.PoStProof, error) {
	return nil, errPostNotBuiltWithCGo
}

func (m *Manager) GenerateWindowPoStWithVanilla(_ context.Context, _ abi.RegisteredPoStProof, _ abi.ActorID, _ abi.PoStRandomness, _ [][]byte, _ int) (proof.PoStProof, error) {
	return proof.PoStProof{}, errPostNotBuiltWithCGo
}
