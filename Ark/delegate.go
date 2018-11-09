package Ark

type Delegate struct {
	Voters  map[string]Voter
	Address string
}

func NewDelegate() *Delegate {
	return &Delegate{Voters: make(map[string]Voter)}
}
