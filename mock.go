package bolt

import (
	"os"

	boltdb "github.com/boltdb/bolt"
)

type DB struct {
	*boltdb.DB
}

type Options struct {
	*boltdb.Options
}

func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	db, err := boltdb.Open(os.DevNull, mode, options.Options)
	if err != nil {
		return nil, err
	}
	return &DB{db}, err
}

func (db *DB) Update(fn func(*Tx) error) error {
	return nil
}

func (db *DB) View(fn func(*Tx) error) error {
	return nil
}

type Bucket struct {
	*boltdb.Bucket
}

func (b *Bucket) Cursor() *Cursor {
	return nil
}

type Cursor struct {
	*boltdb.Cursor
}

type Tx struct {
	*boltdb.Tx
}

func (tx *Tx) Bucket(name []byte) *Bucket {
	return nil
}

func (tx *Tx) Cursor() *Cursor {
	return nil
}
