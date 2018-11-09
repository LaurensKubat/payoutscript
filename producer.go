package payoutscript

import (
	"context"
	"time"
)

// Producer should be implemented by any producer of blocks.
type BlockProducer interface {
	GetBlocks(blockchan chan Block, ctx context.Context)
	GetTotalState() ([]Voter, time.Time)
}
