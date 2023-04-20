package repository

import (
	"sync"
)

type TinyUrlInMemRepository struct {
	mux sync.RWMutex
	m   map[string]string
}

func (r *TinyUrlInMemRepository) Add(fullUrl, tinyUrl string) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.m[fullUrl] = tinyUrl

	return nil
}

func (r *TinyUrlInMemRepository) Get(tinyUrl string) (string, error) {
	var val string
	var ok bool

	r.mux.RLock()
	defer r.mux.RUnlock()

	if val, ok = r.m[tinyUrl]; !ok {
		return "", nil
	}

	return val, nil
}

func (r *TinyUrlInMemRepository) CheckIfFullUrlExists(fullUrl string) (string, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	for tiny, full := range r.m {
		if full == fullUrl {
			return tiny, nil
		}
	}

	return "", nil
}

func (r *TinyUrlInMemRepository) CheckIfTinyUrlExists(tinyUrl string) (bool, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	if _, ok := r.m[tinyUrl]; !ok {
		return false, nil
	}

	return true, nil
}
