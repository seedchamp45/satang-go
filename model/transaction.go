package model

import (
	"math/big"
	"time"
)

// Transaction represents an Ethereum transaction
type Transaction struct {
	Hash      string
	To        string
	Value     *big.Int
	Timestamp time.Time
}
