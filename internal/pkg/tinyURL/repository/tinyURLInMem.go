package repository

import (
	"sync"
)

type TinyUrlInMemoryRepository struct {
	Mux sync.RWMutex
	DB  map[string]string
}

func (r *TinyUrlInMemoryRepository) Add(fullUrl, tinyUrl string) error {
	r.Mux.Lock()
	defer r.Mux.Unlock()

	r.DB[tinyUrl] = fullUrl

	return nil
}

func (r *TinyUrlInMemoryRepository) Get(tinyUrl string) (string, error) {
	var val string
	var ok bool

	r.Mux.RLock()
	defer r.Mux.RUnlock()

	if val, ok = r.DB[tinyUrl]; !ok {
		return "", nil
	}

	return val, nil
}

func (r *TinyUrlInMemoryRepository) CheckIfFullUrlExists(fullUrl string) (string, error) {
	r.Mux.RLock()
	defer r.Mux.RUnlock()

	for tiny, full := range r.DB {
		if full == fullUrl {
			return tiny, nil
		}
	}

	return "", nil
}

func (r *TinyUrlInMemoryRepository) CheckIfTinyUrlExists(tinyUrl string) (bool, error) {
	r.Mux.RLock()
	defer r.Mux.RUnlock()

	if _, ok := r.DB[tinyUrl]; !ok {
		return false, nil
	}

	return true, nil
}
