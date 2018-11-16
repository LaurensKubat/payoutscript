package ark

import (
	"context"
	ark "github.com/ArkEcosystem/go-client/client/two"
	"github.com/LaurensKubat/payoutscript"
	acrypty "github.com/ArkEcosystem/go-crypto/crypto"
)

type API struct {
	delegate  Delegate
	client    *ark.Client
}

const Limit  = 500

func (a *API) GetBlocks(blockchan chan payoutscript.Block, ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:

	}
}

func (a *API) getAllDelegateVotes() ([]ark.Transaction, error) {
	maxPage :=  1
	var transactions []ark.Transaction
	for i := 1; i <= maxPage; i++ {
		res, _, err := a.client.Votes.List(context.Background(), &ark.Pagination{Limit:Limit, Page: i})
		if err != nil {
			return nil, err
		}
		maxPage = int(res.Meta.TotalCount)
		for _, tx := range res.Data {
			if tx.Address == a.delegate.Address {
				transactions = append(transactions, tx)
			}
		}
	}
	return transactions, nil
}

func (a *API) getLatestVoterPayout() {
	maxPage := 1
	payouts := make(map[string]ark.Transaction )
	for i := 1; i <= maxPage; i++ {
		res, _, err := a.client.Wallets.SentTransactions(context.Background(), a.delegate.Address,
			&ark.Pagination{Limit: Limit, Page: i})
		for _, tx := range res.Data {

		}
	}
}