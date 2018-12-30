package payoutscript

type Block struct {
	Height int
	Voters map[VoterAddress]Voter
	Value  int64
}

func NewBlock() *Block {
	return &Block{
		Voters: make(map[VoterAddress]Voter),
	}
}
