package payoutscript

import (
	"time"
)

type Voter struct {
	Address       VoterAddress
	VoteTimestamp time.Time
	Stake         int64
	Percentage    int64
	isVoter       bool
}

type VoterAddress string
