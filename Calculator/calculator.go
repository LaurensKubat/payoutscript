package Calculator

import (
	"github.com/LaurensKubat/payoutscript"
	"github.com/pkg/errors"
	"time"
)

type ShareCalc struct {
	share        float64
	NewBlocks    chan payoutscript.Block
	currentState State
}

type State struct {
	Voters    map[payoutscript.VoterAddress]payoutscript.Voter
	timestamp time.Time
}

const NoBlock = "No new block"
const After = "New state is before current state"

func NewShareCalc(share float64) *ShareCalc {
	return &ShareCalc{
		share: share,
	}
}

func (s *ShareCalc) NextBlock() (payoutscript.Block, error) {
	select {
	case block := <-s.NewBlocks:
		return block, nil
	default:
		return payoutscript.Block{}, errors.New(NoBlock)
	}
}

func (s *ShareCalc) getTotalState(p payoutscript.BlockProducer) error {
	voters, timestamp := p.GetTotalState()
	if s.currentState.timestamp.After(timestamp) {
		return errors.New(After)
	}
	s.currentState.timestamp = timestamp
	for _, voter := range voters {
		s.currentState.Voters[voter.Address] = voter
	}
	return nil
}
