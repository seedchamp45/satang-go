package ethereum

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/seedchamp45/satang-go/model"
	"github.com/seedchamp45/satang-go/pkg/database"
)

// StartTransactionMonitoring starts monitoring incoming and outgoing transactions for the specified addresses
func StartTransactionMonitoring(client *EthereumClient, db *database.PostgreSQLDB, addresses []string) {
	// Create a channel for receiving new block headers
	headers := make(chan *types.Header)

	// Subscribe to new block headers
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("Subscription error:", err)
		case header := <-headers:
			go monitorTransactions(client, db, header.Number, addresses)
		}
	}
}

// monitorTransactions monitors transactions in a block and stores relevant ones in the database
func monitorTransactions(client *EthereumClient, db *database.PostgreSQLDB, blockNumber *big.Int, addresses []string) {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Println("Failed to retrieve block", blockNumber, ":", err)
		return
	}

	for _, tx := range block.Transactions() {
		transaction := model.Transaction{
			Hash:      tx.Hash().Hex(),
			To:        tx.To().Hex(),
			Value:     tx.Value(),
			Timestamp: time.Now().UTC(),
		}

		err := db.StoreTransaction(transaction)
		if err != nil {
			log.Println("Failed to store transaction", tx.Hash().Hex(), ":", err)
		} else {
			log.Println("Stored transaction", tx.Hash().Hex())
		}
	}
}

// containsAddress checks if the given address exists in the list of addresses
func containsAddress(addresses []string, address *common.Address) bool {
	for _, a := range addresses {
		if address != nil && address.Hex() == a {
			return true
		}
	}
	return false
}
