package payoutscript

import (
	"time"
)

type Voter struct {
	Address       VoterAddress
	IsVoter       bool
	VoteTimestamp time.Time
}

type VoterAddress string
