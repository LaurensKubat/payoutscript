package calculator

import (
	"github.com/LaurensKubat/payoutscript"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewShareCalc(t *testing.T) {
	assert.NotPanics(t, func() {
		NewShareCalc(nil)
	})
}

func TestCalculator_calculateBalance(t *testing.T) {

}

type testData struct {
	block      map[payoutscript.VoterAddress]payoutscript.Voter
	amount     int64
	share      int64
	shouldPass bool
}

func TestCalculator_verifyBalance(t *testing.T) {
	// we set up some testdata
	data := []testData{
		{
			block: map[payoutscript.VoterAddress]payoutscript.Voter{
				"aaa": {
					Address: "aaa",
					Stake:   300,
				},
				"assadsd": {
					Address: "assadsd",
					Stake:   2,
				},
				"asdasda": {
					Address: "asdasda",
					Stake:   1231412,
				},
			},
			amount:     2000,
			share:      95,
			shouldPass: true,
		},
		{
			block: map[payoutscript.VoterAddress]payoutscript.Voter{
				"aaasdaa": {
					Address: "aaa",
					Stake:   123,
				},
				"asbdahwdbw": {
					Address: "asbdahwdbw",
					Stake:   433,
				},
				"rpmyjhbldgkn": {
					Address: "rpmyjhbldgkn",
					Stake:   44567,
				},
			},
			amount:     2000000,
			share:      95.0,
			shouldPass: true,
		},
	}
	for _, dataset := range data {
		s := NewShareCalc(nil)
		s.updateState(dataset.block, time.Now())
		ok := s.verifyBalance(dataset.amount)
		assert.True(t, ok)
	}
}

func TestCalculator_CalculateBlock(t *testing.T) {
	data := []testData{
		{
			block: map[payoutscript.VoterAddress]payoutscript.Voter{
				"aaa": {
					Address: "aaa",
					Stake:   300,
				},
				"assadsd": {
					Address: "assadsd",
					Stake:   2,
				},
				"asdasda": {
					Address: "asdasda",
					Stake:   1231412,
				},
			},
			amount:     2000,
			share:      95.0,
			shouldPass: true,
		},
		{
			block: map[payoutscript.VoterAddress]payoutscript.Voter{
				"aaasdaa": {
					Address: "aaasdaa",
					Stake:   123,
				},
				"asbdahwdbw": {
					Address: "asbdahwdbw",
					Stake:   433,
				},
				"rpmyjhbldgkn": {
					Address: "rpmyjhbldgkn",
					Stake:   44567,
				},
			},
			amount:     2000000,
			share:      95.0,
			shouldPass: true,
		},
	}
	for _, dataset := range data {
		s := NewShareCalc(nil)
		s.updateState(dataset.block, time.Now())
		res, err := s.calculateBalance(dataset.amount)
		assert.NoError(t, err)
		//we check that all calculated addresses exist in our dataset
		for address, _ := range res {
			assert.Equal(t, address, dataset.block[address].Address)
		}
	}
}

func TestCalculator_NextBlock(t *testing.T) {
	// we set up some data
	data := []payoutscript.Block{
		{
			Timestamp: time.Now(),
			Voters:    nil,
			Value:     2,
		},
		{
			Timestamp: time.Now(),
			Voters:    nil,
			Value:     1,
		},
	}

	blocks := make(chan payoutscript.Block, 10)
	for _, data := range data {
		blocks <- data
	}
	s := NewShareCalc(blocks)
	for i := range data {
		block, err := s.NextBlock()
		assert.NoError(t, err)
		assert.Equal(t, block, data[i])
	}
	_, err := s.NextBlock()
	assert.Error(t, err)
}
