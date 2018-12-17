package payoutscript

import (
	"time"
)

type Voter struct {
	Address       VoterAddress
	PubKey        string
	VoteTimestamp time.Time
	Balance       int64
	Percentage    int64
	isVoter       bool
}

type VoterAddress string
