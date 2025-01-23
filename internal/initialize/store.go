package initialize

import (
	"lime/internal/global"
	"lime/pkg/store"
	"path/filepath"
)

func InitStore() {
	path := filepath.Join(global.ROOT_PATH, "data", "store")
	store, err := store.NewPebbleStore(path)
	if err != nil {
		panic(err)
	}

	global.STORE = store
}
