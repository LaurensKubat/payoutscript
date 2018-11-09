package Ark

import "github.com/LaurensKubat/payoutscript"

type Delegate struct {
	Voters  map[string]payoutscript.Voter
	Address string
}

func NewDelegate() *Delegate {
	return &Delegate{Voters: make(map[string]payoutscript.Voter)}
}
