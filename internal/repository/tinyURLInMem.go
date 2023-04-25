package repository

import (
	"sync"
)

type TinyUrlInMemoryRepository struct {
	mux sync.RWMutex
	db  map[string]string
}

func InitTinyUrlInMemoryRepository() *TinyUrlInMemoryRepository {
	return &TinyUrlInMemoryRepository{mux: sync.RWMutex{}, db: make(map[string]string)}
}

func (r *TinyUrlInMemoryRepository) Add(fullUrl, tinyUrl string) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.db[tinyUrl] = fullUrl

	return nil
}

func (r *TinyUrlInMemoryRepository) Get(tinyUrl string) (string, error) {
	var val string
	var ok bool

	r.mux.RLock()
	defer r.mux.RUnlock()

	if val, ok = r.db[tinyUrl]; !ok {
		return "", nil
	}

	return val, nil
}

func (r *TinyUrlInMemoryRepository) CheckIfFullUrlExists(fullUrl string) (string, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	for tiny, full := range r.db {
		if full == fullUrl {
			return tiny, nil
		}
	}

	return "", nil
}

func (r *TinyUrlInMemoryRepository) CheckIfTinyUrlExists(tinyUrl string) (bool, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	if _, ok := r.db[tinyUrl]; !ok {
		return false, nil
	}

	return true, nil
}
