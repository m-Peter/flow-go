package badger

import (
	"fmt"

	"github.com/dgraph-io/badger/v2"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/storage/badger/operation"
)

type Collections struct {
	db *badger.DB
}

func NewCollections(db *badger.DB) *Collections {
	c := Collections{
		db: db,
	}
	return &c
}

func (c *Collections) ByFingerprint(hash flow.Fingerprint) (*flow.Collection, error) {
	var collection flow.Collection

	err := c.db.View(func(tx *badger.Txn) error {
		return operation.RetrieveCollection(hash, &collection)(tx)
	})
	if err != nil {
		return nil, fmt.Errorf("could not retrieve collection: %w", err)
	}

	return &collection, nil
}

func (c *Collections) TransactionsByFingerprint(hash flow.Fingerprint) ([]*flow.Transaction, error) {
	var (
		collection flow.Collection
		transactions []*flow.Transaction
	)

	err := c.db.View(func(tx *badger.Txn) error {
		err := operation.RetrieveCollection(hash, &collection)(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve collection: %w", err)
		}

		for _, txHash := range collection.Transactions {
			var transaction flow.Transaction
			err = operation.RetrieveTransaction(txHash, &transaction)(tx)
			if err != nil {
				return fmt.Errorf("could not retrieve transaction: %w", err)
			}

			transactions = append(transactions, &transaction)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (c *Collections) Insert(collection *flow.Collection) error {
	return c.db.Update(func (tx *badger.Txn) error {
		err := operation.InsertCollection(collection.Fingerprint(), collection)(tx)
		if err != nil {
			return fmt.Errorf("could not insert collection: %w", err)
		}
		return nil
	})
}

func (c *Collections) Remove(hash flow.Fingerprint) error {
	return c.db.Update(func(tx *badger.Txn) error {
		err := operation.RemoveCollection(hash)(tx)
		if err != nil {
			return fmt.Errorf("could not remove collection: %w", err)
		}
		return nil
	})
}

func (c *Collections) InsertGuarantee(gc *flow.GuaranteedCollection) error {
	return c.db.Update(func(tx *badger.Txn) error {
		err := operation.InsertGuaranteedCollection(gc.Fingerprint(), gc)(tx)
		if err != nil {
			return fmt.Errorf("could not insert guaranteed collection: %w", err)
		}
		return nil
	})
}

func (c *Collections) RetrieveGuarantee(hash flow.Fingerprint) (*flow.GuaranteedCollection, error) {
	var gc flow.GuaranteedCollection

	err := c.db.View(func(tx *badger.Txn) error {
		return operation.RetrieveGuaranteedCollection(hash, &gc)(tx)
	})
	if err != nil {
		return nil, fmt.Errorf("could not retrieve guaranteed collection: %w", err)
	}

	return &gc, nil
}
