package Ark

import (
	"context"
	ark "github.com/ArkEcosystem/go-client/client/two"
	arkcrypto "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/payoutscript"
)

type API struct {
	share float64
	delegate Delegate
	client *ark.Client
}

func (a *API) GetCurrentVoterState() error {
	curpage := 1
	max := 1
	for ; curpage <= max; curpage++ {
		resp, _, err := a.client.Delegates.Voters(context.Background(),
			a.delegate.Address, &ark.Pagination{Page:curpage, Limit: 500})
		max = int(resp.Meta.PageCount)
		if err != nil {
			return err
		}

		for _, wallet := range resp.Data {
			if _, registered := a.delegate.Voters[wallet.Address]; !registered {
				a.delegate.Voters[wallet.Address] = payoutscript.Voter{
					Address:       wallet.Address,
					IsVoter:       true,
					VoteTimestamp: nil,
				}
			}
		}
	}
	return nil
}

func (a *API) GetPastVotersStates() error {
	for address, voter := range a.delegate.Voters {
		curpage := 1
		resp, _, err := a.client.Wallets.Transactions(context.Background(), address,
			&ark.Pagination{Page: curpage, Limit: 500})
		if err != nil {
			return err
		}
		for _, transaction := range resp.Data {
			if transaction.Type == arkcrypto.TRANSACTION_TYPES.Vote {
				if voter.VoteTimestamp ==  {// add
					voter.VoteTimestamp = transaction.Timestamp
				}
			}
		}
	}
}
