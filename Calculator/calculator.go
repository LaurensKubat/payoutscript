package Calculator

import (
	"github.com/LaurensKubat/payoutscript/payoutscript"
	"github.com/pkg/errors"
)

type ShareCalc struct {
	share     float64
	NewBlocks chan payoutscript.Block
}

const NoBlock = "No new block"

func NewShareCalc(share float64) *ShareCalc {
	return &ShareCalc{
		share: share,
	}
}

func (s ShareCalc) NextBlock() (payoutscript.Block, error) {
	select {
	case block := <-s.NewBlocks:
		return block, nil
	default:
		return payoutscript.Block{}, errors.New(NoBlock)
	}
}
