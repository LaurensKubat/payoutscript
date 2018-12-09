package arkdb

import (
	"database/sql"
	"github.com/LaurensKubat/payoutscript"
)
import _ "github.com/lib/pq"

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
}

// get all transactions by anyone who ever voted on the delegate
func (d *DataScraper) getTransactions(voteSerialized string) (Transactions, error) {
	rows, err := d.db.Query(`
		SELECT transactions.id, transactions.amount, blocks.height, transactions.recipient_id,
		transactions.sender_public_key, transanctions.serialized, transactions.type, transactions.fee, transactions.block_id
		FROM transactions 
		  INNER JOIN blocks ON transactions.block_id = blocks.id 
			WHERE transactions.recipient_id IN (
			  SELECT transactions.recipient_id 
			  FROM transactions 
			  WHERE transactions.serialized LIKE '%?%'
			  )
	`, voteSerialized)

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

type ForgedBlocks []ForgedBlock

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
		blocks = append(blocks, block)
	}
	return blocks, nil
}

type voter struct {
	address     string
	pubKey      string
	voteBalance int64
	vote        string
}

type Voters []voter

// get all voters who are voters at the moment of running the payout script
func (d *DataScraper) getCurrentVoters(pubKey string) (Voters, error) {
	rows, err := d.db.Query(`
	SELECT wallets.address, wallets.public_key, wallets.vote_balance, wallets.vote
	FROM wallets
	WHERE wallets.vote = ?`, pubKey)

	if err != nil {
		return nil, err
	}
	voters := Voters{}
	for rows.Next() {
		voter := voter{}
		err := rows.Scan(&voter.address, &voter.pubKey, &voter.voteBalance, &voter.vote)
		if err != nil {
			return nil, err
		}
		voters = append(voters, voter)
	}
	return voters, nil

}

// TODO get a list of all transactions ordered by voter and if those voters have forged any blocks, look up those blocks
// and add their reward to the total balance of that voter
type state struct {
	transaction     int
	amount          int64
	height          int
	recipientID     string
	senderPublicKey string
	serialized      []byte
	transactionType int
	transactionFee  int64
	blockID         int64
	producedBlocks  int64
}

type States []state

func (d *DataScraper) getTotalState(serialized string) (States, error) {
	rows, err := d.db.Query(`
	SELECT transactions.id, transactions.amount, blocks.height, transactions.recipient_id,
		transactions.sender_public_key, transanctions.serialized, transactions.type, transactions.fee,
	       transactions.block_id, wallets.produced_blocks
		FROM transactions 
		  INNER JOIN blocks 
		    ON transactions.block_id = blocks.id
		  INNER JOIN wallets 
		    ON transactions.recipient_id = wallets.address
		WHERE transactions.recipient_id IN (
		  SELECT transactions.recipient_id 
			FROM transactions 
			  WHERE transactions.serialized LIKE '%?%'
			  ) `, serialized)

	if err != nil {
		return nil, err
	}

	var states States
	for rows.Next() {
		state := state{}
		err := rows.Scan(&state.transaction, &state.amount, &state.height, &state.recipientID, &state.senderPublicKey,
			&state.serialized, &state.transactionType, &state.transactionFee, &state.blockID, &state.producedBlocks)
		if err != nil {
			return nil, err
		}
		states = append(states, state)
	}
	return states, nil
}

func (d *DataScraper) StatesToBlocks(states States) []payoutscript.Block {
	for _, state := range states {

	}
}
