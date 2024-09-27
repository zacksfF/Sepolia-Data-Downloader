package db

import (
	"encoding/binary"
	"encoding/json"
	"log"
	bchain "sepolia_doanloader/b_chain"

	"github.com/syndtr/goleveldb/leveldb"
)

type BlockStore struct {
	DB *leveldb.DB
}

// Initialize LevelDB
func InitDB(path string) *BlockStore {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatalf("Failed to open LevelDB: %v", err)
	}

	return &BlockStore{DB: db}
}

// StoreBlock saves a Block to the LevelDB keyed by its Index
func (store *BlockStore) StoreBlock(block bchain.Block) error {
	// Serialize block to JSON
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}

	// Use block index as key
	key := bchain.Int64Bytes(int64(block.Index))

	// Store block data in LevelDB
	err = store.DB.Put(key, data, nil)
	return err
}

// GetBlock retrieves a block from LevelDB by its index
func (store *BlockStore) GetBlock(index int64) (bchain.Block, error) {
	var block bchain.Block

	key := bchain.Int64Bytes(index)
	data, err := store.DB.Get(key, nil)
	if err != nil {
		return block, err
	}

	// Deserialize JSON to Block struct
	err = json.Unmarshal(data, &block)
	return block, err
}

// Convert int64 to byte slice for LevelDB key
func Int64ToBytes(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

type L1Info struct {
	BlockTime   uint64
	ParentHash  string
	L1InfoRoot  []byte
	BlockNumber uint64
	Index       int
}

// StoreL1Info saves the L1 info to LevelDB
func (store *BlockStore) StoreL1Info(info L1Info) error {
	// Serialize L1Info to JSON
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}

	// Use incrementing index as the key
	key := Int64ToBytes(int64(info.Index))

	// Store the L1 info in LevelDB
	err = store.DB.Put(key, data, nil)
	return err
}
