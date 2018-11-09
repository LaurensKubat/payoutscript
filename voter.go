package payoutscript

import (
	"time"
)

type Voter struct {
	Address       VoterAdress
	IsVoter       bool
	VoteTimestamp time.Time
}

type VoterAdress string
