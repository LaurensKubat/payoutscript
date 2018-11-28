package ark

import (
	"github.com/ArkEcosystem/go-crypto/crypto"
	"github.com/LaurensKubat/payoutscript"
)

type TxBuilder struct {
}

// BuildBlock builds one block (represented as a map) into an array of transactions. Use BuildBlockIntoBatch for
// multiple blocks.
func (t *TxBuilder) BuildBlock(voters map[payoutscript.VoterAddress]int64, passphrase1 string, passphrase2 string,
	vendorfield string) []*crypto.Transaction {
	Txs := []*crypto.Transaction{}
	for address, amount := range voters {
		tx := crypto.BuildTransfer(string(address), uint64(amount), vendorfield, passphrase1, passphrase2)
		Txs = append(Txs, tx)
	}
	return Txs
}

// concatenate multiple calculated blocks (represented as only a map of addresses and amounts) into one map
func (t *TxBuilder) concatBlocks(maps []map[payoutscript.VoterAddress]int64) map[payoutscript.VoterAddress]int64 {
	concat := map[payoutscript.VoterAddress]int64{}
	for _, voters := range maps {
		for address, amount := range voters {
			concat[address] += amount
		}
	}
	return concat
}

// BuildBlocksIntoBatch makes a batch of multiple blocks (represented as an array of maps) into multiple transactions.
// Transactions to the same address are added together. make the 2nd passphrase an empty string if you do not want to
// use the 2nd passphrase
func (t *TxBuilder) BuildBlocksIntoBatch(voters []map[payoutscript.VoterAddress]int64, passphrase1 string,
	passphrase2 string, vendorfield string) []*crypto.Transaction {
	concatvoters := t.concatBlocks(voters)
	Txs := []*crypto.Transaction{}
	for address, amount := range concatvoters {
		tx := crypto.BuildTransfer(string(address), uint64(amount), vendorfield, passphrase1, passphrase2)
		Txs = append(Txs, tx)
	}
	return Txs
}
