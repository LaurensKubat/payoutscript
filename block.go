package payoutscript

type Block struct {
	Timestamp Timestamp
	NewVotes  map[VoterAddress]Voter
	UnVotes   map[VoterAddress]Voter
}

type Timestamp struct {
	Epoch int64
	Unix  int64
	Human string
}

func NewBlock() *Block {
	return &Block{
		NewVotes: make(map[VoterAddress]Voter),
		UnVotes:  make(map[VoterAddress]Voter),
	}
}

func (b *Block) NewVoter() {

}

func (b *Block) NewUnvote() {

}
