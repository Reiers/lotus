//go:build !cgo
// +build !cgo

// Stub implementations of the two CGo-dependent methods on *Local so
// that storage/paths satisfies the Store interface under CGO_ENABLED=0.
// These methods are not exercised by PDP-only consumers (Curio Core);
// any caller that invokes them under !cgo gets a clear error pointing
// at the missing build configuration. Behavior parity with the
// equivalent split in github.com/Reiers/curio's lib/paths.

package paths

import (
	"context"
	"errors"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/storage/sealer/storiface"
)

var errPathsNotBuiltWithCGo = errors.New("paths: this method requires building with CGO_ENABLED=1 (filecoin-ffi linkage)")

// GenerateSingleVanillaProof is a CGo-gated stub under !cgo.
// Returns errPathsNotBuiltWithCGo at runtime; never panics. Curio Core
// PDP flows never call this; sealing PoSt does.
func (st *Local) GenerateSingleVanillaProof(ctx context.Context, minerID abi.ActorID, si storiface.PostSectorChallenge, ppt abi.RegisteredPoStProof) ([]byte, error) {
	_ = ctx
	_ = minerID
	_ = si
	_ = ppt
	return nil, errPathsNotBuiltWithCGo
}

// GeneratePoRepVanillaProof is a CGo-gated stub under !cgo.
// Returns errPathsNotBuiltWithCGo at runtime; never panics. Used only
// by sealing PoRep, which is out of scope for Curio Core.
func (st *Local) GeneratePoRepVanillaProof(ctx context.Context, sr storiface.SectorRef, sealed, unsealed cid.Cid, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness) ([]byte, error) {
	_ = ctx
	_ = sr
	_ = sealed
	_ = unsealed
	_ = ticket
	_ = seed
	return nil, errPathsNotBuiltWithCGo
}
