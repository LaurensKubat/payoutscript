package payoutscript

import ark "github.com/ArkEcosystem/go-client/client/two"

type Voter struct {
	Address       string
	IsVoter       bool
	VoteTimestamp ark.Timestamp
}
