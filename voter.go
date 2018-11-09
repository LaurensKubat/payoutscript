package payoutscript

import (
	"time"
)

type Voter struct {
	Address       string
	IsVoter       bool
	VoteTimestamp time.Time
}
