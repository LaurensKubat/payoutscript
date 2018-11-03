package payoutscript

import ark "github.com/ArkEcosystem/go-client/client/two"

type ShareCalc struct {
	share float64
	delegateID string
	client *ark.Client
}

func NewShareCalc(delegateID string, share float64) *ShareCalc {
	return &ShareCalc{
		share:share,
		delegateID:delegateID,
		client:ark.NewClient(nil),
	}
}

func (s ShareCalc) CalculateLatestBlock (delegateID string) {

}

func (s ShareCalc)