# Sepolia Data Downloader for L2 Project

This project is a **Sepolia Data Downloader** tool designed to efficiently query logs from the Ethereum Sepolia testnet as part of an L2 blockchain project. It extracts specific data from a smart contract, stores L1 info, and saves the information in a key-value store (LevelDB). The downloader focuses on querying and storing:

- **L1 Info Root Data**
- **Block Time**
- **Parent Hash**

## Table of Contents

- [Project Overview](#project-overview)
- [Key Features](#key-features)
- [Technologies Used](#technologies-used)
- [Installation and Setup](#installation-and-setup)
- [Usage](#usage)
- [Structure](#structure)
- [License](#license)

## Project Overview

The purpose of this project is to efficiently download and store data from the Sepolia Ethereum testnet, with a focus on a particular smart contract and topic. The downloaded data is stored in a LevelDB instance for later use, making it useful for an L2 project requiring blockchain data retrieval, indexing, and storage.

## Key Features

- Query Sepolia Ethereum testnet logs using specific contract and topic.
- Retrieve block information, including block time, parent hash, and L1 info root data.
- Store the extracted information in **LevelDB**, using an incrementing index for each event.
- Handles multiple logs within a block efficiently by using log indexes.
- Ensures the validity of blockchain data through cryptographic verification.

## Technologies Used

- **Go**: The main programming language used for building this application.
- **Ethereum Go Client**: For querying and interacting with the Ethereum blockchain.
- **LevelDB**: A key-value store to efficiently store and retrieve blockchain data.
- **SHA-256 Cryptography**: Used to hash and validate blocks.
- **Infura**: Ethereum node provider used to connect to the Sepolia testnet.

## Installation and Setup

### Prerequisites

Before you can use this tool, make sure you have the following installed:

- Go (Golang) v1.22 or higher
- LevelDB
- An Infura account (to get an Ethereum node endpoint for Sepolia)

### Clone the Repository

```bash
git clone https://github.com/zacksfF/Sepolia-Data-Downloader.git
cd Sepolia-Data-Downloader
```

### Install Dependencies

To install the required Go packages, run:

```bash
go mod tidy
```

### Configure the Ethereum Client

Ensure you have access to the Sepolia network through Infura:

- Replace the Infura project ID in the `client/ethereum.go` file:

```go
client, err := ethclient.Dial("https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID")
```

### Build the Application

You can build the project by running:

```bash
go build -o Sepolia-Data-Downloader
```

## Usage

After building the application, run it as follows:

```bash
./Sepolia-Data-Downloader
```

The tool will connect to the Sepolia testnet, query for logs from the specified contract and topic, and store the L1 info root data, block time, and parent hash in LevelDB.

### Command-line Arguments

You can modify the block range and other parameters directly in the `eth/client.go` file for now:

```go
startBlock := big.NewInt(1000000) // Replace with the actual start block
endBlock := big.NewInt(1005000)   // Replace with the actual end block
```

## Structure

This project consists of the following main components:

- **`l1chain.go`**: Defines the blockchain block structure and validation mechanisms.
- **`eth.go`**: Interacts with the Sepolia Ethereum network, queries for logs, and extracts L1 info.
- **`store.go`**: Manages the storage and retrieval of L1 info and block data using LevelDB.
- **`main.go`**: The entry point to the application, initializes everything and runs the process.

### Example Code:

```go
package main

import (
	"log"
	"sepolia_doanloader/client"
	"sepolia_doanloader/db"
)

func main() {
	// Initialize LevelDB
	store := db.InitDB("./blockdata")

	// Query Sepolia and store L1 info
	if err := client.QuerySepolia(store); err != nil {
		log.Fatalf("Error querying Sepolia: %v", err)
	}

	log.Println("Data downloaded and stored successfully!")
}
```

### Data Storage:

The **L1 info** extracted from the Sepolia logs is stored in LevelDB as a serialized JSON object with an incrementing index for each event log. Each block of data includes:

- **BlockTime**: The timestamp of the block.
- **ParentHash**: The parent block hash.
- **L1InfoRoot**: The root data from the log event.
- **BlockNumber**: The number of the block.
- **Index**: A unique index for each log within the block.



### Testing a qucik implemt of (layer 1) Blockchain Implementation

#### Running Tests
To test the blockchain, you can run:

```bash
cd b_chain
go test
```


####  Output
```json
{
  "Index": 0,
  "Timestamp": "2024-09-27T17:09:53.362402+01:00",
  "Data": null,
  "PreviousHash": null,
  "Hash": "r5jygeLj/oGREtEsIGFopz/r4gkd8VWVoiB/6nvys5g="
}
```
