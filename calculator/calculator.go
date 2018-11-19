package calculator

import (
	"github.com/LaurensKubat/payoutscript"
	"github.com/pkg/errors"
	"time"
)

// calculator checks that no blocks before its current state are calculated because the calculation is dependant on it's
// current state.
type Calculator struct {
	NewBlocks    chan payoutscript.Block
	currentState State
	delegate     payoutscript.Delegate
}

type State struct {
	Voters    map[payoutscript.VoterAddress]payoutscript.Voter
	timestamp time.Time
}

const (
	NoBlock             = "No new block"
	After               = "New state is before current state"
	InsufficientBalance = "insufficient balance for the calculation"
)

func NewShareCalc(NewBlocks chan payoutscript.Block) *Calculator {
	return &Calculator{
		NewBlocks: NewBlocks,
	}
}

func (s *Calculator) NextBlock() (payoutscript.Block, error) {
	select {
	case block := <-s.NewBlocks:
		return block, nil
	default:
		return payoutscript.Block{}, errors.New(NoBlock)
	}
}

func (s *Calculator) getTotalState(p payoutscript.BlockProducer) error {
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

func (s *Calculator) CalculateBlock(block payoutscript.Block) (map[payoutscript.VoterAddress]int64, error) {
	err := s.updateState(block.Voters, block.Timestamp)
	if err != nil {
		return nil, err
	}
	if !s.verifyBalance(block.Value) {
		return nil, errors.New(InsufficientBalance)
	}
	payouts, err := s.calculateBalance(block.Value)
	if err != nil {
		return nil, err
	}
	return payouts, nil
}

func (s *Calculator) verifyBalance(amount int64) bool {
	var totalVoterStake int64
	for _, voter := range s.currentState.Voters {
		totalVoterStake += voter.Stake
	}
	for _, voter := range s.currentState.Voters {
		amount -= amount * voter.Percentage * (voter.Stake / totalVoterStake)
		if amount < 0 {
			return false
		}
	}
	return true
}

func (s *Calculator) calculateBalance(amount int64) (map[payoutscript.VoterAddress]int64, error) {
	var totalVoterStake int64
	rewardPerVoter := make(map[payoutscript.VoterAddress]int64)
	for _, voter := range s.currentState.Voters {
		totalVoterStake += voter.Stake
	}
	for _, voter := range s.currentState.Voters {
		toPay := amount * voter.Percentage * (voter.Stake / totalVoterStake)
		amount -= amount * voter.Percentage * (voter.Stake / totalVoterStake)
		if amount < 0 {
			return nil, errors.New("Calculate Balance got a negative amount, while verify balance passed")
		}
		rewardPerVoter[voter.Address] = toPay
	}
	return rewardPerVoter, nil
}

func (s *Calculator) updateState(NewVoters map[payoutscript.VoterAddress]payoutscript.Voter, ts time.Time) error {
	if s.currentState.timestamp.After(ts) {
		return errors.New(After)
	}
	s.currentState.timestamp = ts
	s.currentState.Voters = NewVoters
	return nil
}
