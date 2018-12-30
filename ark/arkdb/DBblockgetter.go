package arkdb

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	arkcrypto "github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/LaurensKubat/payoutscript"
	"strings"
)

type DataScraper struct {
	db *sql.DB
}

func NewDataScraper(db *sql.DB) *DataScraper {
	return &DataScraper{db: db}
}

type Transactions []transaction

type transaction struct {
	Id              string
	Amount          int64
	Height          int
	RecipientID     string
	SenderPubKey    string
	Serialized      []byte
	Transactiontype int
	Fee             int64
	BlockID         string
	ProducedBlocks  int
}

type stringslice struct {
	Slice []string
	Valid bool
}

func (s stringslice) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return "'" + strings.Join(s.Slice, "', '") + "'", nil
}

// get all transactions by anyone who ever voted on the delegate
// the qry is made in a weird way because recipientIDs and pubKeys have a variable length
func (d *DataScraper) GetEvents(recipientIDs []interface{}, pubKeys []interface{}) (Transactions, error) {
	qry := `SELECT transactions.id, transactions.amount, blocks.height, transactions.recipient_id,
			transactions.sender_public_key, transactions.serialized, transactions.type, transactions.fee,
			transactions.block_id
		FROM transactions 
		  INNER JOIN blocks 
		    ON transactions.block_id = blocks.id 
		WHERE transactions.recipient_id IN ('a')
		  OR transactions.sender_public_key IN  )
ORDER BY blocks.height ASC;`
	fmt.Println(qry)
	rows, err := d.db.Query(qry, recipientIDs)

	if err != nil {
		return Transactions{}, err
	}

	transactions := Transactions{}
	for rows.Next() {
		transaction := transaction{}
		err := rows.Scan(&transaction.Id, &transaction.Amount, &transaction.Height, &transaction.RecipientID,
			&transaction.SenderPubKey, &transaction.Serialized, &transaction.Transactiontype, &transaction.Fee,
			&transaction.BlockID)
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
func (d *DataScraper) GetForgedBlocks(pubKey string) (ForgedBlocks, error) {
	rows, err := d.db.Query(`
	SELECT blocks.height, blocks.total_fee, blocks.reward, generator_public_key
	FROM blocks
	WHERE generator_public_key = $1
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
func (d *DataScraper) GetVoters(delegatePubKey string) (Voters, error) {
	serialized, err := d.getSerializedVoters()
	if err != nil {
		fmt.Println("cannot get voters")
		return nil, err
	}
	voters := make(map[payoutscript.VoterAddress]Voter)
	for i := 0; i < len(serialized); i++ {
		deser := arkcrypto.DeserializeTransaction(serialized[i])
		for _, vote := range deser.Asset.Votes {
			if strings.Contains(vote, delegatePubKey) {
				voters[payoutscript.VoterAddress(deser.RecipientId)] = Voter{
					Address: payoutscript.VoterAddress(deser.RecipientId),
					PubKey:  deser.SenderPublicKey,
				}
			}
		}
	}
	return voters, nil
}
func (d *DataScraper) getSerializedVoters() ([]string, error) {
	rows, err := d.db.Query(`
	SELECT encode(transactions.serialized::bytea, 'hex')
	FROM transactions
	WHERE transactions.type = $1`, arkcrypto.TRANSACTION_TYPES.Vote)

	if err != nil {
		return nil, err
	}
	votes := []string{}
	for rows.Next() {
		serialized := ""
		err := rows.Scan(&serialized)
		if err != nil {
			return nil, err
		}
		votes = append(votes, serialized)
	}
	return votes, nil
}
