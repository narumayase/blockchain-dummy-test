package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockchain_IsValid_true(t *testing.T) {
	blockchain := NewBlockchain()
	blockchain.AddBlock(
		"{\"from\":\"viet\"," +
			"\"to\":\"david\"," +
			"\"amount\":\"100\"}")

	blockchain.AddBlock(
		"{\"from\":\"juan\"," +
			"\"to\":\"pepe\"," +
			"\"amount\":\"150\"}")

	for _, c := range blockchain.Chain {
		fmt.Println(c.String())
	}
	fmt.Printf("isValid() = %t\n", blockchain.IsValid())

	assert.True(t, blockchain.IsValid(), "has to be true if no one edited the blockchain")
}

func TestBlockchain_IsValid_false(t *testing.T) {
	blockchain := NewBlockchain()
	blockchain.AddBlock(
		"{\"from\":\"viet\"," +
			"\"to\":\"david\"," +
			"\"amount\":\"100\"}")

	blockchain.AddBlock(
		"{\"from\":\"juan\"," +
			"\"to\":\"pepe\"," +
			"\"amount\":\"150\"}")

	for _, c := range blockchain.Chain {
		fmt.Println(c.String())
	}

	blockchain.Chain[1].Data = "{\"from\":\"juan\"," +
		"\"to\":\"pepe\"," +
		"\"amount\":\"15000000000\"}"

	fmt.Printf("isValid() = %t\n", blockchain.IsValid())

	assert.False(t, blockchain.IsValid(), "has to be false if no the blockchain is corrupted")
}
