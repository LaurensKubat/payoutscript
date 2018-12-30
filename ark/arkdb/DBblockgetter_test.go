package arkdb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func newDataScraper() (*DataScraper, error) {
	connStr := os.Getenv("CONNSTR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return NewDataScraper(db), nil
}

func TestNewDataScraper(t *testing.T) {
	assert.NotPanics(t, func() {
		connStr := os.Getenv("CONNSTR")
		db, err := sql.Open("postgres", connStr)
		assert.NoError(t, err)
		NewDataScraper(db)
	})
}

func TestDataScraper_GetEvents(t *testing.T) {
	t.Parallel()
	d, err := newDataScraper()
	require.NoError(t, err)
	var addresses, pubkeys []string
	v, err := d.GetVoters("0218b77efb312810c9a549e2cc658330fcc07f554d465673e08fa304fa59e67a0a")
	for address, voter := range v {
		addresses = append(addresses, string(address))
		pubkeys = append(pubkeys, voter.PubKey)
	}
	res, err := d.GetEvents(addresses, pubkeys)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	fmt.Println(res)
}

func TestDataScraper_GetForgedBlocks(t *testing.T) {
	t.Parallel()
	d, err := newDataScraper()
	require.NoError(t, err)
	f, err := d.GetForgedBlocks("0218b77efb312810c9a549e2cc658330fcc07f554d465673e08fa304fa59e67a0a")
	assert.NoError(t, err)
	assert.NotNil(t, f)
}

func TestDataScraper_GetVoters(t *testing.T) {
	t.Parallel()
	d, err := newDataScraper()
	require.NoError(t, err)
	v, err := d.GetVoters("0218b77efb312810c9a549e2cc658330fcc07f554d465673e08fa304fa59e67a0a")
	assert.NoError(t, err)
	assert.NotNil(t, v)
}
