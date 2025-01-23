package store

import (
	"fmt"
	"os"

	pebbledb "github.com/cockroachdb/pebble"
)

type PebbleStore struct {
	db   *pebbledb.DB
	mode *pebbledb.WriteOptions
}

func NewPebbleStore(path string) (Store, error) {
	// 确保目录存在
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("创建目录失败: %v", err)
	}

	store := &PebbleStore{}
	store.mode = pebbledb.Sync

	// 配置数据库选项
	opts := &pebbledb.Options{
		BytesPerSync:        1 << 20, // 1MB
		MaxOpenFiles:        1000,
		MemTableSize:        64 << 20, // 64MB
		MaxConcurrentCompactions: func() int { return 4 },
	}

	db, err := pebbledb.Open(path, opts)
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %v", err)
	}

	store.db = db
	return store, nil
}

func (p *PebbleStore) Get(key string) ([]byte, error) {
	data, _, err := p.db.Get([]byte(key))
	return data, err
}

func (p *PebbleStore) Set(key string, value []byte) error {
	return p.db.Set([]byte(key), value, p.mode)
}

func (p *PebbleStore) Del(key string) error {
	return p.db.Delete([]byte(key), p.mode)
}

func (p *PebbleStore) Close() error {
	return p.db.Close()
}

func (p *PebbleStore) ForEach(prefix string, fn func(key string, value []byte) error) error {
	keyUpperBound := func(b []byte) []byte {
		end := make([]byte, len(b))
		copy(end, b)
		for i := len(end) - 1; i >= 0; i-- {
			end[i] = end[i] + 1
			if end[i] != 0 {
				return end[:i+1]
			}
		}
		return nil // no upper-bound
	}

	prefixIterOptions := func(prefix []byte) *pebbledb.IterOptions {
		return &pebbledb.IterOptions{
			LowerBound: prefix,
			UpperBound: keyUpperBound(prefix),
		}
	}

	iter, err := p.db.NewIter(prefixIterOptions([]byte(prefix)))
	if err != nil {
		return err
	}

	defer iter.Close()

	for iter.First(); iter.Valid(); iter.Next() {
		err := fn(string(iter.Key()), iter.Value())
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PebbleStore) Count(prefix string) (int, error) {
	count := 0
	err := p.ForEach(prefix, func(key string, value []byte) error {
		count++
		return nil
	})
	return count, err
}

func (p *PebbleStore) Clear(prefix string) error {
	return p.ForEach(prefix, func(key string, value []byte) error {
		return p.db.Delete([]byte(key), p.mode)
	})
}
