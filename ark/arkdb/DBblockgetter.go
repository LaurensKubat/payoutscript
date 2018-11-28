package arkdb

import (
	"database/sql"
)
import _ "github.com/lib/pq"

type DataScraper struct {
	db *sql.DB
}

func NewDataScraper(db *sql.DB) *DataScraper {
	return &DataScraper{db: db}
}

type Transactions struct {
	transactions []transaction
}

type transaction struct {
	id              string
	amount          int64
	timestamp       int
	recipientID     string
	senderPubKey    string
	serialized      []byte
	transactiontype int
	fee             int64
	blockID         string
}

// TODO add
func (d *DataScraper) getTransactions() (sql.Result, error) {
	res, err := d.db.Exec(`
		SELECT transactions.id, transactions.amount, blocks.timestamp, transactions.recipient_id,
		transactions.sender_public_key, transanctions.serialized, transactions.type, transactions.fee, transactions.block_id
		FROM transactions 
		  INNER JOIN blocks ON transactions.block_id = blocks.id 
			WHERE transactions.type = 3 AND //
	`)

	return res, err
}

func (d *DataScraper) getLastOutTransactions(pubKey string) (sql.Result, error) {
	res, err := d.db.Exec(`
	SELECT 
	transaction.recipient_id, transactions.timestamp, blocks.height, transactions.id, transaction.amount, transaction.fee
	FROM transaction
	  INNER JOIN blocks
 	  ON (blocks.id = transaction.block_id),
	(SELECT 
	 MAX(transaction.timestamp) AS max_timestamp,
	 transaction.recipient_id
	 FROM transactions AS ts
	 WHERE transactions.sender_public_key = ?
	 GROUP BY transactions.sender_public_key) AS maxresults

	WHERE transactions.recipient_id = maxresults.recipient_id
	AND transactions.timestamp = maxresults.max_timestamp;`, pubKey)
	return res, err
}

func (d *DataScraper)


