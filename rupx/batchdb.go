package rupx

import (
	"bytes"
	"encoding/hex"
	"github.com/rupayaproject/rupaya/common"
	"github.com/rupayaproject/rupaya/rupx/rupx_state"
	"github.com/globalsign/mgo"
	"sync"

	"github.com/rupayaproject/rupaya/ethdb"
	"github.com/rupayaproject/rupaya/log"
	lru "github.com/hashicorp/golang-lru"
)

const (
	defaultCacheLimit = 1024
)

type BatchItem struct {
	Value interface{}
}

type BatchDatabase struct {
	db         *ethdb.LDBDatabase
	emptyKey   []byte
	cacheItems *lru.Cache // Cache for reading
	lock       sync.RWMutex
	cacheLimit int
	Debug      bool
}

// NewBatchDatabase use rlp as encoding
func NewBatchDatabase(datadir string, cacheLimit int) *BatchDatabase {
	return NewBatchDatabaseWithEncode(datadir, cacheLimit)
}

// batchdatabase is a fast cache db to retrieve in-mem object
func NewBatchDatabaseWithEncode(datadir string, cacheLimit int) *BatchDatabase {
	db, err := ethdb.NewLDBDatabase(datadir, 128, 1024)
	if err != nil {
		log.Error("Can't create new DB", "error", err)
		return nil
	}
	itemCacheLimit := defaultCacheLimit
	if cacheLimit > 0 {
		itemCacheLimit = cacheLimit
	}

	cacheItems, _ := lru.New(itemCacheLimit)

	batchDB := &BatchDatabase{
		db:         db,
		cacheItems: cacheItems,
		emptyKey:   rupx_state.EmptyKey(), // pre alloc for comparison
		cacheLimit: itemCacheLimit,
	}

	return batchDB

}

func (db *BatchDatabase) IsEmptyKey(key []byte) bool {
	return key == nil || len(key) == 0 || bytes.Equal(key, db.emptyKey)
}

func (db *BatchDatabase) getCacheKey(key []byte) string {
	return hex.EncodeToString(key)
}

func (db *BatchDatabase) HasObject(hash common.Hash) (bool, error) {
	// for mongodb only
	return false, nil
}

func (db *BatchDatabase) GetObject(hash common.Hash, val interface{}) (interface{}, error) {
	// for mongodb only
	return nil, nil
}

func (db *BatchDatabase) PutObject(hash common.Hash, val interface{}) error {
	// for mongodb only
	return nil
}

func (db *BatchDatabase) DeleteObject(hash common.Hash) error {
	// for mongodb only
	return nil
}

func (db *BatchDatabase) Put(key []byte, val []byte) error {
	return db.db.Put(key, val)
}

func (db *BatchDatabase) Delete(key []byte) error {
	return db.db.Delete(key)
}

func (db *BatchDatabase) Has(key []byte) (bool, error) {
	return db.db.Has(key)
}

func (db *BatchDatabase) Get(key []byte) ([]byte, error) {
	return db.db.Get(key)
}

func (db *BatchDatabase) Close() {
	db.db.Close()
}

func (db *BatchDatabase) NewBatch() ethdb.Batch {
	return db.db.NewBatch()
}

func (db *BatchDatabase) DeleteTradeByTxHash(txhash common.Hash) {
}

func (db *BatchDatabase) GetOrderByTxHash(txhash common.Hash) []*rupx_state.OrderItem {
	return []*rupx_state.OrderItem{}
}

func (db *BatchDatabase) GetListOrderByHashes(hashes []string) []*rupx_state.OrderItem {
	return []*rupx_state.OrderItem{}
}

func (db *BatchDatabase) InitBulk() *mgo.Session {
	return nil
}

func (db *BatchDatabase) CommitBulk() error {
	return nil
}
