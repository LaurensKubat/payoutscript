package ark

import (
	"context"
	ark "github.com/ArkEcosystem/go-client/client/two"
	"github.com/LaurensKubat/payoutscript"
)


// current pagination limit as defined in https://github.com/ArkEcosystem/core/tree/develop/packages/core-api
const Limit = 100
//the timestamp of DD, which decides the payout percentage for is
const DDOriginTS = int32(16247647)

type BlockBuilder struct {
	client   *ark.Client
}

func (b *BlockBuilder) getVoterPercentage(ts ark.Timestamp) int64 {
	if ts.Epoch < DDOriginTS {
		return 96
	}
	return 95
}

func (b *BlockBuilder)getCurrentVoters(delegatePubKey string) (map[payoutscript.VoterAddress]payoutscript.Voter, error) {
	Voters :=  make(map[payoutscript.VoterAddress]payoutscript.Voter)
	total := 2
	for page := 1; page <= int(total); page++ {
		res, _, err := b.client.Delegates.Voters(context.Background(), delegatePubKey,
			&ark.Pagination{Limit: Limit, Page:page})
		if err != nil {
			return nil, err
		}
		total = int(res.Meta.PageCount)
		for _, voter := range res.Data {
			newVoter := payoutscript.Voter{
				Address: payoutscript.VoterAddress(voter.Address),
				PubKey: voter.PublicKey,
				Balance: voter.Balance
			}
			Voters[payoutscript.VoterAddress(voter.Address)] = newVoter
		}
	}
	return Voters, nil
}