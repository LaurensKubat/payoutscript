package payoutscript

type producer interface {
	GetChannel() chan Block
}
