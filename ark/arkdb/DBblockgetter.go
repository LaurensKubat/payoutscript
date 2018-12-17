package arkdb

import (
	"database/sql"
	arkcrypto "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/LaurensKubat/payoutscript"
)

type DataScraper struct {
	db *sql.DB
}

func NewDataScraper(db *sql.DB) *DataScraper {
	return &DataScraper{db: db}
}

type Transactions []transaction

type transaction struct {
	id              string
	amount          int64
	height          int
	recipientID     string
	senderPubKey    string
	serialized      []byte
	transactiontype int
	fee             int64
	blockID         string
	producedBlocks  int
}

// get all transactions by anyone who ever voted on the delegate
func (d *DataScraper) getEvents(recipientID []string, pubKey []string) (Transactions, error) {
	rows, err := d.db.Query(`
		SELECT transactions.id, transactions.amount, blocks.height, transactions.recipient_id,
			transactions.sender_public_key, transanctions.serialized, transactions.type, transactions.fee,
			transactions.block_id, wallets.produced_blocks
		FROM transactions 
		  INNER JOIN blocks 
		    ON transactions.block_id = blocks.id 
		  INNER JOIN wallets 
		    on transactions.recipient_id = wallets.address
			WHERE transactions.recipient_id IN (
			  SELECT transactions.recipient_id 
			  FROM transactions 
			  WHERE transactions.recipient_id IN ?
				OR transactions.sender_public_key IN ?
			  )
		ORDER BY blocks.height ASC
	`, recipientID, pubKey)

	if err != nil {
		return Transactions{}, err
	}

	transactions := Transactions{}
	for rows.Next() {
		transaction := transaction{}
		err := rows.Scan(&transaction.id, &transaction.amount, &transaction.height, &transaction.recipientID,
			&transaction.senderPubKey, &transaction.serialized, &transaction.transactiontype, &transaction.fee,
			&transaction.blockID)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

type ForgedBlocks map[int64]ForgedBlock

type ForgedBlock struct {
	Height   int64
	TotalFee int64
	Reward   int64
	PubKey   string
}

// Select all blocks forged by the delegate and the corresponding fees and rewards
func (d *DataScraper) getForgedBlocks(pubKey string) (ForgedBlocks, error) {
	rows, err := d.db.Query(`
	SELECT blocks.height, blocks.total_fee, blocks.reward, generator_public_key
	FROM blocks
	WHERE generator_public_key = ?
	ORDER BY blocks.height DESC`, pubKey)

	if err != nil {
		return ForgedBlocks{}, err
	}
	blocks := ForgedBlocks{}
	for rows.Next() {
		block := ForgedBlock{}
		err := rows.Scan(&block.Height, &block.TotalFee, &block.Reward, &block.PubKey)
		if err != nil {
			return nil, err
		}
		blocks[block.Height] = block
	}
	return blocks, nil
}

type Voter struct {
	Address        payoutscript.VoterAddress
	PubKey         string
	ProducedBlocks int
}
type Voters map[payoutscript.VoterAddress]Voter

// get all voters who are voters at the moment of running the payout script
func (d *DataScraper) getVoters(delegateName string) (Voters, error) {
	serialized, err := d.getSerializedVoters()
	if err != nil {
		return nil, err
	}
	voters := make(map[payoutscript.VoterAddress]Voter)
	for i := 0; i < len(serialized); i++ {
		deser := arkcrypto.DeserializeTransaction(string(serialized[i]))
		if deser.Asset.Delegate.Username == delegateName {
			voters[payoutscript.VoterAddress(deser.RecipientId)] = Voter{
				Address: payoutscript.VoterAddress(deser.RecipientId),
				PubKey:  deser.SenderPublicKey,
			}
		}
	}
	return voters, nil
}
func (d *DataScraper) getSerializedVoters() ([][]byte, error) {
	rows, err := d.db.Query(`
	SELECT transactions.serialized
	FROM transactions
	WHERE transactions.type = ?`, int(arkcrypto.TRANSACTION_TYPES.Vote))

	if err != nil {
		return nil, err
	}
	votes := [][]byte{}
	for rows.Next() {
		serialized := []byte{}
		err := rows.Scan(&serialized)
		if err != nil {
			return nil, err
		}
		votes = append(votes, serialized)
	}
	return votes, nil
}

// TODO get a list of all transactions ordered by voter and if those voters have forged any blocks, look up those blocks
// and add their reward to the total balance of that voter

// statesToBlocks creates blocks from the states we query above, it then iterates backwards over all transactions,
// keeping the voter databases, until it reaches the block height where the last payout was made
func (d *DataScraper) statesToBlocks(delegateName string, delegatePubKey string) ([]payoutscript.Block, error) {
	voters, err := d.getVoters(delegateName)
	if err != nil {
		return nil, err
	}

	//make a slice of the voter addresses and a slice of the public keys so we can fetch all transactions
	voterAddresses := []string{}
	publicKeys := []string{}
	for _, voter := range voters {
		voterAddresses = append(voterAddresses, string(voter.Address))
		publicKeys = append(publicKeys, voter.PubKey)
	}
	events, err := d.getEvents(voterAddresses, publicKeys)
	forgedBlocks, err := d.getForgedBlocks(delegatePubKey)
	// we use events[0].height because the events are ordered by blocks.height in the SQL query
	for i := events[0].height; i >= 0; i-- {
		block := payoutscript.Block{}
		if forged, exists := forgedBlocks[int64(events[i].height)]; exists {
			block.Value = forged.Reward
		}
	}
}
