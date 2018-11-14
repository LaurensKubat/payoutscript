package payoutscript

import (
	"time"
)

type Voter struct {
	Address       VoterAddress
	VoteTimestamp time.Time
	Stake         float64
}

type VoterAddress string
