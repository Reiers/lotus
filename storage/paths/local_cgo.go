//go:build cgo
// +build cgo

// CGo-gated implementations of the two filecoin-ffi-bound methods on
// *Local. Split out so storage/paths can build under CGO_ENABLED=0
// for PDP-only consumers (Curio Core). Byte-identical to the upstream
// implementation when CGo is enabled; only the location moved.

package paths

import (
	"context"
	"time"

	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/proof"

	"github.com/filecoin-project/lotus/lib/result"
	"github.com/filecoin-project/lotus/storage/sealer/storiface"
)

func (st *Local) GenerateSingleVanillaProof(ctx context.Context, minerID abi.ActorID, si storiface.PostSectorChallenge, ppt abi.RegisteredPoStProof) ([]byte, error) {
	sr := storiface.SectorRef{
		ID: abi.SectorID{
			Miner:  minerID,
			Number: si.SectorNumber,
		},
		ProofType: si.SealProof,
	}

	var cache, sealed, cacheID, sealedID string

	if si.Update {
		src, si, err := st.AcquireSector(ctx, sr, storiface.FTUpdate|storiface.FTUpdateCache, storiface.FTNone, storiface.PathStorage, storiface.AcquireMove)
		if err != nil {
			return nil, xerrors.Errorf("acquire sector: %w", err)
		}
		cache, sealed = src.UpdateCache, src.Update
		cacheID, sealedID = si.UpdateCache, si.Update
	} else {
		src, si, err := st.AcquireSector(ctx, sr, storiface.FTSealed|storiface.FTCache, storiface.FTNone, storiface.PathStorage, storiface.AcquireMove)
		if err != nil {
			return nil, xerrors.Errorf("acquire sector: %w", err)
		}
		cache, sealed = src.Cache, src.Sealed
		cacheID, sealedID = si.Cache, si.Sealed
	}

	if sealed == "" || cache == "" {
		return nil, errPathNotFound
	}

	psi := ffi.PrivateSectorInfo{
		SectorInfo: proof.SectorInfo{
			SealProof:    si.SealProof,
			SectorNumber: si.SectorNumber,
			SealedCID:    si.SealedCID,
		},
		CacheDirPath:     cache,
		PoStProofType:    ppt,
		SealedSectorPath: sealed,
	}

	start := time.Now()

	resCh := make(chan result.Result[[]byte], 1)
	go func() {
		resCh <- result.Wrap(ffi.GenerateSingleVanillaProof(psi, si.Challenge))
	}()

	select {
	case r := <-resCh:
		return r.Unwrap()
	case <-ctx.Done():
		log.Errorw("failed to generate valilla PoSt proof before context cancellation", "err", ctx.Err(), "duration", time.Since(start), "cache-id", cacheID, "sealed-id", sealedID, "cache", cache, "sealed", sealed)

		// this will leave the GenerateSingleVanillaProof goroutine hanging, but that's still less bad than failing PoSt
		return nil, xerrors.Errorf("failed to generate vanilla proof before context cancellation: %w", ctx.Err())
	}
}

func (st *Local) GeneratePoRepVanillaProof(ctx context.Context, sr storiface.SectorRef, sealed, unsealed cid.Cid, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness) ([]byte, error) {
	src, _, err := st.AcquireSector(ctx, sr, storiface.FTSealed|storiface.FTCache, storiface.FTNone, storiface.PathStorage, storiface.AcquireMove)
	if err != nil {
		return nil, xerrors.Errorf("acquire sector: %w", err)
	}

	if src.Sealed == "" || src.Cache == "" {
		return nil, errPathNotFound
	}

	ssize, err := sr.ProofType.SectorSize()
	if err != nil {
		return nil, xerrors.Errorf("getting sector size: %w", err)
	}

	secPiece := []abi.PieceInfo{{
		Size:     abi.PaddedPieceSize(ssize),
		PieceCID: unsealed,
	}}

	return ffi.SealCommitPhase1(sr.ProofType, sealed, unsealed, src.Cache, src.Sealed, sr.ID.Number, sr.ID.Miner, ticket, seed, secPiece)
}
