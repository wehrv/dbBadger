package dbBadger // github.com/wehrv/dbBadger

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

var db *badger.DB

func init() {
	var err error
	db, err = badger.Open(badger.DefaultOptions("data"))
	if err != nil {
		log.Fatal(err)
	}
}

func Pull(key []byte) ([]byte, error) {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	itm, err := txn.Get(key)
	var val []byte
	if err == nil {
		val, err = itm.ValueCopy(val)
	}
	return val, err
}

func Push(key, val []byte) error {
	txn := db.NewTransaction(true)
	defer txn.Discard()
	err := txn.Set(key, val)
	if err == nil {
		err = txn.Commit()
	}
	return err
}

func Drop(key []byte) error {
	txn := db.NewTransaction(true)
	defer txn.Discard()
	err := txn.Delete(key)
	if err == nil {
		err = txn.Commit()
	}
	return err
}

func Done() error {
	return db.Close()
}
