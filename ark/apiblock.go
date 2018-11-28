package ark

import (
	"context"
	ark "github.com/ArkEcosystem/go-client/client/two"
	"github.com/LaurensKubat/payoutscript"
	"github.com/LaurensKubat/payoutscript/ark/arkutils"
)

type API struct {
	delegate Delegate
	client   *ark.Client
}

// current pagination limit as defined in https://github.com/ArkEcosystem/core/tree/develop/packages/core-api
const Limit = 100
const DDOriginTS = int32(16247647)

func (a *API) GetBlocks(blockchan chan payoutscript.Block, ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:

	}
}

//TODO make sure all api calls are done simultaniously
func (a *API) getAllDelegateVotes() ([]ark.Transaction, error) {
	maxPage := 1
	var transactions []ark.Transaction
	for i := 1; i <= maxPage; i++ {
		res, _, err := a.client.Votes.List(context.Background(), &ark.Pagination{Limit: Limit, Page: i})
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

//TODO look at the indexing of the transactions and see if i can incorporate the time in the key for efficient lookup

//getVoterPayouts gets all sent transactions from the delegate wallet
func (a *API) getVoterPayouts() (map[string]ark.Transaction, error) {
	maxPage := 1
	payouts := make(map[string]ark.Transaction)
	for i := 1; i <= maxPage; i++ {
		res, _, err := a.client.Wallets.SentTransactions(context.Background(), a.delegate.Address,
			&ark.Pagination{Limit: Limit, Page: i})
		if err != nil {
			return nil, err
		}
		for _, tx := range res.Data {
			payouts[tx.Address] = tx
		}
	}
	return payouts, nil
}

func (a *API) getBlocks() ([]ark.Block, error) {
	var blocks []ark.Block
	maxPage := 1
	for i := 1; i <= maxPage; i++ {
		res, _, err := a.client.Delegates.Blocks(context.Background(), a.delegate.Address,
			&ark.Pagination{Limit: Limit, Page: i})
		if err != nil {
			return nil, err
		}
		for _, block := range res.Data {
			//TODO change this to immediatly convert to my own block type
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

func (a *API) getVoters() (map[int32]payoutscript.Voter, error) {
	var voters map[int32]payoutscript.Voter
	votes, err := a.getAllDelegateVotes()
	if err != nil {
		return nil, err
	}
	for _, vote := range votes {
		voter := payoutscript.Voter{
			Address:    payoutscript.VoterAddress(vote.Sender),
			Stake:      vote.Amount,
			Percentage: a.getVoterPercentage(vote.Timestamp),
		}
		voter.VoteTimestamp = arkutils.ParseArkTs(vote.Timestamp)
		voters[vote.Timestamp.Unix] = voter
	}
	return voters, nil
}

func (a *API) getVoterPercentage(ts ark.Timestamp) int64 {
	if ts.Epoch < DDOriginTS {
		return 96
	}
	return 95
}
