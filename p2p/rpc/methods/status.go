package methods

import (
	"encoding/hex"
	"fmt"
	"github.com/protolambda/rumor/p2p/rpc/reqresp"
)

type Status struct {
	ForkDigest      ForkDigest
	FinalizedRoot   Root
	FinalizedEpoch  Epoch
	HeadRoot        Root
	HeadSlot        Slot
}

func (s *Status) Data() map[string]interface{} {
	return map[string]interface{}{
		"fork_digest": hex.EncodeToString(s.ForkDigest[:]),
		"finalized_root": hex.EncodeToString(s.FinalizedRoot[:]),
		"finalized_epoch": s.FinalizedEpoch,
		"head_root": hex.EncodeToString(s.HeadRoot[:]),
		"head_slot": s.HeadSlot,
	}
}

func (s *Status) String() string {
	return fmt.Sprintf("Status(fork_digest: %x, finalized_root: %x, finalized_epoch: %d, head_root: %x, head_slot: %d)",
		s.ForkDigest, s.FinalizedRoot, s.FinalizedEpoch, s.HeadRoot, s.HeadSlot)
}

var StatusRPCv1 = reqresp.RPCMethod{
	Protocol:                  "/eth2/beacon_chain/req/status/1/ssz",
	RequestCodec:              reqresp.NewSSZCodec((*Status)(nil)),
	ResponseChunkCodec:        reqresp.NewSSZCodec((*Status)(nil)),
	DefaultResponseChunkCount: 1,
}
