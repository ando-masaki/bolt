package bolt

import "os"

type DB struct {
}

type Options struct {
}

func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	if _, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, mode); err != nil {
		return nil, err
	}
	return &DB{}, nil
}

func (db *DB) Update(fn func(*Tx) error) error {
	return nil
}

func (db *DB) View(fn func(*Tx) error) error {
	return nil
}

func (db *DB) Close() error {
	return nil
}

type Bucket struct {
	FillPercent float64
}

func (b *Bucket) Cursor() *Cursor {
	return &Cursor{}
}

func (b *Bucket) Put(key []byte, value []byte) error {
	return nil
}

func (b *Bucket) Get(key []byte) []byte {
	return nil
}

type Cursor struct {
}

type Tx struct {
}

func (tx *Tx) Bucket(name []byte) *Bucket {
	return &Bucket{}
}

func (tx *Tx) Cursor() *Cursor {
	return &Cursor{}
}

func (tx *Tx) CreateBucketIfNotExists(name []byte) (*Bucket, error) {
	return &Bucket{}, nil
}
