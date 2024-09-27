package client

import (
	"context"
	"log"
	"math/big"
	"sepolia_doanloader/db"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddress = "0x761d53b47334bee6612c0bd1467fb881435375b2"                         //The Contract Address
	topic           = "0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766" //the Hash Topic
)

func QuerySepolia(store *db.BlockStore) error {
	// Connect to Sepolia testnet
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/02e307e2b52941bd986f0ff21c09236a")
	if err != nil {
		log.Fatalf("Failed to connect to Sepolia: %v", err)
	}
	defer client.Close()

	// Define the block range to query (e.g., last 5000 blocks)
	startBlock := big.NewInt(1000000) // Replace with the actual start block
	endBlock := big.NewInt(1005000)   // Replace with the actual end block

	// Create a filter query
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{
			common.HexToAddress(contractAddress),
		},
		Topics: [][]common.Hash{
			{common.HexToHash(topic)},
		},
	}

	// Fetch logs matching the query
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to fetch logs: %v", err)
	}

	// Process each log
	for _, vLog := range logs {
		block, err := client.BlockByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			log.Printf("Failed to retrieve block for hash %s: %v", vLog.BlockHash.Hex(), err)
			continue
		}

		// Extract L1 info (block time, parent hash, and log data)
		l1Info := db.L1Info{
			BlockTime:   block.Time(),
			ParentHash:  block.ParentHash().Hex(),
			L1InfoRoot:  vLog.Data, // This would be the data for the L1 info root
			BlockNumber: vLog.BlockNumber,
			Index:       int(vLog.Index), // Unique index for each log within a block
		}

		// Store the extracted data in LevelDB
		err = store.StoreL1Info(l1Info)
		if err != nil {
			log.Printf("Failed to store L1 info: %v", err)
		}
	}

	return nil
}
