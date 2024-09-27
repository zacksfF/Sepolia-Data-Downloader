package main

import (
	"fmt"
	"log"
	bchain "sepolia_doanloader/b_chain"
	"sepolia_doanloader/client"
	"sepolia_doanloader/db"
)

func main() {
	// Initialize LevelDB for storing blocks
	store := db.InitDB("blockstore.db")
	defer store.DB.Close()

	fmt.Println("LevelDB initialized and ready.")

	// Initialize blockchain
	chain := bchain.NewChain()
	fmt.Println("Blockchain initialized.")

	// Query Sepolia testnet and store the logs in the blockchain
	err := client.QuerySepolia(store)
	if err != nil {
		log.Fatalf("Error querying Sepolia: %v", err)
	}
	fmt.Println("Sepolia data downloaded and stored.")

	// Verify blockchain integrity
	isValid := chain.Verify()
	if !isValid {
		log.Fatalf("Blockchain verification failed!")
	}
	fmt.Println("Blockchain verification successful.")
}
