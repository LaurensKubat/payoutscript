package Calculator

import (
	ark "github.com/ArkEcosystem/go-client/client/two"
)

type ShareCalc struct {
	share    float64
	delegate Delegate
	client   *ark.Client
}

func NewShareCalc(delegate Delegate, share float64) *ShareCalc {
	return &ShareCalc{
		share:    share,
		delegate: delegate,
		client:   ark.NewClient(nil),
	}
}
