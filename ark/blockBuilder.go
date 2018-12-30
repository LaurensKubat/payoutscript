package ark

import (
	"context"
	ark "github.com/ArkEcosystem/go-client/client/two"
	arkcrypto "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/LaurensKubat/payoutscript"
	"github.com/LaurensKubat/payoutscript/ark/arkdb"
)

// current pagination limit as defined in https://github.com/ArkEcosystem/core/tree/develop/packages/core-api
const Limit = 100

//the timestamp of DD, which decides the payout percentage for is
const DDOriginTS = int32(16247647)

type BlockBuilder struct {
	client *ark.Client
	db     *arkdb.DataScraper
}

func (b *BlockBuilder) getVoterPercentage(ts ark.Timestamp) int64 {
	if ts.Epoch < DDOriginTS {
		return 96
	}
	return 95
}

// TODO get a list of all transactions ordered by voter and if those voters have forged any blocks, look up those blocks
// and add their reward to the total balance of that voter

// statesToBlocks creates blocks from the states we query above, it then iterates backwards over all transactions,
// keeping the voter databases, until it reaches the block height where the last payout was made

func (b *BlockBuilder) ParseBlocks(delegateName string, delegatePubKey string) ([]payoutscript.Block, error) {
	voters, err := b.db.GetVoters(delegateName)
	if err != nil {
		return nil, err
	}

	events, err := b.getEvents(voters)
	if err != nil {
		return nil, err
	}
	forgedBlocks, err := b.db.GetForgedBlocks(delegatePubKey)
	if err != nil {
		return nil, err
	}
	// we use events[0].height because the events are ordered by blocks.height DESC in the SQL query
	curHeight := events[0].Height
	curVoters, err := b.getCurrentVoters(delegatePubKey)
	if err != nil {
		return nil, err
	}
	var blocks []payoutscript.Block
	block := payoutscript.Block{}
	for i := len(events) - 1; i >= 0; i-- {
		// if the next tx is in a different block, finish this block
		if events[i].Height != curHeight {
			block.Voters = curVoters
			block.Height = curHeight
			blocks = append(blocks, block)
			block = payoutscript.Block{}

		}
		// check if we forged the block
		if forged, exists := forgedBlocks[int64(events[i].Height)]; exists {
			block.Value = forged.Reward
		}
		if events[i].Transactiontype == int(arkcrypto.TRANSACTION_TYPES.Vote) {
			deser := arkcrypto.DeserializeTransaction(string(events[i].Serialized))
		}
	}
}

func (b *BlockBuilder) getEvents(voters arkdb.Voters) (arkdb.Transactions, error) {
	//make a slice of the voter addresses and a slice of the public keys so we can fetch all transactions
	voterAddresses := []string{}
	publicKeys := []string{}
	for _, voter := range voters {
		voterAddresses = append(voterAddresses, string(voter.Address))
		publicKeys = append(publicKeys, voter.PubKey)
	}
	return b.db.GetEvents(voterAddresses, publicKeys)
}

// getCurrentVoters gets all current voters of the delegate using the ark api
func (b *BlockBuilder) getCurrentVoters(delegatePubKey string) (map[payoutscript.VoterAddress]payoutscript.Voter, error) {
	Voters := make(map[payoutscript.VoterAddress]payoutscript.Voter)
	total := 1
	for page := 1; page <= int(total); page++ {
		res, _, err := b.client.Delegates.Voters(context.Background(), delegatePubKey,
			&ark.Pagination{Limit: Limit, Page: page})
		if err != nil {
			return nil, err
		}
		total = int(res.Meta.PageCount)
		for _, voter := range res.Data {
			newVoter := payoutscript.Voter{
				Address: payoutscript.VoterAddress(voter.Address),
				PubKey:  voter.PublicKey,
				Balance: voter.Balance,
			}
			Voters[payoutscript.VoterAddress(voter.Address)] = newVoter
		}
	}
	return Voters, nil
}
