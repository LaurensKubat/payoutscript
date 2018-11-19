package payoutscript

import "time"

type Block struct {
	Timestamp time.Time
	Voters    map[VoterAddress]Voter
	Value     int64
}

func NewBlock() *Block {
	return &Block{
		Voters: make(map[VoterAddress]Voter),
	}
}
