package payoutscript

import "context"

// Producer should be implemented by any producer of blocks.
type BlockProducer interface {
	GetBlocks(blockchan chan Block, ctx context.Context)
}
