package payoutscript

type Block struct {
	Timestamp Timestamp
	NewVotes  map[VoterAdress]Voter
	UnVotes   map[VoterAdress]Voter
}

type Timestamp struct {
	Epoch int64
	Unix  int64
	Human string
}

func NewBlock() *Block {
	return &Block{
		NewVotes: make(map[VoterAdress]Voter),
		UnVotes:  make(map[VoterAdress]Voter),
	}
}

func (b *Block) NewVoter() {

}

func (b *Block) NewUnvote() {

}
