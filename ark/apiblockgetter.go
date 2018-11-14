package ark

import (
	"context"
	ark "github.com/ArkEcosystem/go-client/client/two"
	"github.com/LaurensKubat/payoutscript"
)

type API struct {
	delegate Delegate
	client   *ark.Client
}

func (a *API) GetBlocks(blockchan chan payoutscript.Block, ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:

	}
}
