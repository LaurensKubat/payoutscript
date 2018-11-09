package payoutscript

import ()

type Block struct {
	Timestamp Timestamp
	NewVotes  map[string]Voter
	UnVotes   map[string]Voter
}

type Timestamp struct {
	Epoch int64
	Unix  int64
	Human string
}

func NewBlock() *Block {
	return &Block{
		NewVotes: make(map[string]Voter),
		UnVotes:  make(map[string]Voter),
	}
}

func (b *Block) AddVoter() {

}
