package bchain

import (
	"encoding/base64"
	"testing"
	"time"
)

func TestNewBlock(t *testing.T) {
	// Create a new block
	chain := NewChain()

	// Validate index
	if chain == nil || (*chain)[0].Index != 0 {
		t.Errorf("Expected genesis block with index 0, got %d", (*chain)[0].Index)
	}

	// Check timestamp is not zero
	if (*chain)[0].TimesTamp.IsZero() {
		t.Errorf("Expected non-zero timestamp, got zero")
	}

	// Check hash is non-empty
	if len((*chain)[0].Hash) == 0 {
		t.Errorf("Expected non-empty hash, got empty hash")
	}

	// Output the genesis block
	t.Logf("Genesis Block: %+v", *chain)
}

func TestAddBlock(t *testing.T) {
	chain := NewChain()

	// Write new data to the chain
	data := []byte("new block data")
	_, err := chain.Write(data)
	if err != nil {
		t.Errorf("Error writing block: %v", err)
	}

	// Check the length of the chain is 2 (genesis + new block)
	if len(*chain) != 2 {
		t.Errorf("Expected chain length 2, got %d", len(*chain))
	}

	// Check the latest block data
	if len((*chain)[1].Data) == 0 {
		t.Errorf("Expected non-empty data, got empty data")
	}

	// Output the added block
	t.Logf("Added Block: %+v", (*chain)[1])
}

func TestVerifyChain(t *testing.T) {
	chain := NewChain()

	// Add two blocks
	chain.Write([]byte("block1"))
	chain.Write([]byte("block2"))

	// Verify the chain
	if !chain.Verify() {
		t.Errorf("Chain verification failed, blocks are invalid")
	} else {
		t.Log("Chain verification passed")
	}
}

func TestGenesisBlockOutput(t *testing.T) {
	chain := NewChain()
	genesis := (*chain)[0]

	expectedTimestamp := "2024-09-07T17:46:28.008676+01:00" // Adjust this based on the exact timestamp

	// Output test block in the expected format
	output := map[string]interface{}{
		"Index":        genesis.Index,
		"Timestamp":    genesis.TimesTamp.Format(time.RFC3339Nano),
		"Data":         nil,
		"PreviousHash": nil,
		"Hash":         base64.StdEncoding.EncodeToString(genesis.Hash),
	}

	expectedOutput := map[string]interface{}{
		"Index":        0,
		"Timestamp":    expectedTimestamp,
		"Data":         nil,
		"PreviousHash": nil,
		"Hash":         base64.StdEncoding.EncodeToString(genesis.Hash),
	}

	// Compare the output
	if output["Index"] != expectedOutput["Index"] || output["Timestamp"] != expectedOutput["Timestamp"] || output["Hash"] != expectedOutput["Hash"] {
		t.Errorf("Genesis block output mismatch. Expected %+v, got %+v", expectedOutput, output)
	}

	t.Logf("Genesis Block Output: %+v", output)
}
