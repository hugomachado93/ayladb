package main

import (
	"fmt"
	"sync"
)

var storage *KeyValueStore

type KeyValueStore struct {
	data map[string]string
	mu   sync.RWMutex
}

// Singleton: There should be only one KeyValueStore
func NewKeyValueStore() *KeyValueStore {
	if storage == nil {
		storage = &KeyValueStore{data: make(map[string]string)}
	}

	return storage
}

func (kvs *KeyValueStore) Set(key string, value string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()

	kvs.data[key] = value
}

func (kvs *KeyValueStore) Get(key string) (string, error) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	value, ok := kvs.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}

	return value, nil
}
