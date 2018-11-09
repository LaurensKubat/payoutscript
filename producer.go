package payoutscript

// Producer should be implemented by any producer of blocks.
type Producer interface {
	GetChannel() chan Block
}
