package bchain

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"time"
)

var (
	// ErrInvalidBlock is an error returned if an invalid block is added to
	// a block chain where the new block's has and index are invalid
	ErrInvalidBlock = errors.New("error: invalid block")
)

// We need to define the Block | A Block represent a single block in a chain of blocks(Blockchain)
type Block struct {
	Index       int
	TimesTamp   time.Time
	Data        []byte
	PreviosHash []byte
	Hash        []byte
}

// we need to define the newBlock and returns a new empty `Block`
func newBlock() Block {
	n := Block{}
	n.TimesTamp = time.Now()
	n.Hash = hashBlock(n)

	return n
}

func hashBlock(b Block) []byte {
	h := sha256.New()

	h.Write(Int64Bytes(int64(b.Index)))
	h.Write(Int64Bytes(b.TimesTamp.Unix()))
	h.Write(b.Data)
	h.Write(b.PreviosHash)

	return h.Sum(nil)
}

// isValidate validate the current block ('b') with previos block in chain
func (b Block) isValidate(o Block) bool {
	if (bytes.Compare(b.Hash, hashBlock(b)) != 0) ||
		(b.Index != (o.Index + 1)) || (bytes.Compare(b.PreviosHash, o.Hash) != 0) {
		return false
	}
	return true
}

// Generate creates a new block fron the current block which is assumed to be the
// last block in the chain.
func (b Block) Generate(data []byte) Block {
	n := Block{
		Index:       b.Index + 1,
		TimesTamp:   time.Now(),
		Data:        make([]byte, data[0]),
		PreviosHash: b.Hash,
	}
	return n
}

// Chain is a slice of blocks to form the "block chain". Each block is
// connected to the previous block cryptographically by each block's hash
// being in corporated into the next block in the chain. The larger the chain
// grows the higher the integrity of the chain and the more difficult it is to
// temper with or modify previous blocks in the chain.
type Chain []Block

// NewChain creates a new "Block chain" `Chain ` with an initial block already
// created called the "genesis block"
func NewChain() *Chain {
	return &Chain{newBlock()}
}

// Add adds a new block (`block`) to the chain verifying its validity
// If the Block is invalid an error is returned otherwise the block is appended
// to the block chain.
func (c *Chain) Add(block Block) error {
	preBlock := (*c)[len(*c)-1]
	if !block.isValidate(preBlock) {
		return ErrInvalidBlock
	}
	*c = append(*c, block)
	return nil
}

// Write creates a new block with the given data (`data`) and appends it to the
// block chain. This implements the `io.Writer` interface so you can treat the
// block chain as a valid Writer.
func (c *Chain) Write(data []byte) (int, error) {
	prevBlock := (*c)[len(*c)-1]
	block := prevBlock.Generate(data)
	if err := c.Add(block); err != nil {
		return 0, ErrInvalidBlock
	}
	return len(data), nil
}

// Verify verifies the cryptographic hashes of every block inthe chain
// ensuring all blocks are valid and their integrity in tact
func (c *Chain) Verify() bool {
	prevBlock := (*c)[0]
	for _, block := range (*c)[1:] {
		if !block.isValidate(prevBlock) {
			return false
		}
		prevBlock = block
	}
	return true
}
