package store

import (
	"fmt"
	"sync"

)


type KeyValueStore interface {
	Set(string, string)
	Get(string) (string, error)
}

type inmemorykeyvaluestore struct {
	internalMap map[string]string
	mutex sync.RWMutex
}

func NewInMemoryKeyValueStore() KeyValueStore {
	return &inmemorykeyvaluestore{
		internalMap: map[string]string{},
		mutex: sync.RWMutex{},
	}
}


func (kvStore *inmemorykeyvaluestore) Set(key, value string) {
	kvStore.mutex.Lock()
	kvStore.internalMap[key] = value
	kvStore.mutex.Unlock()
}

func (kvStore *inmemorykeyvaluestore) Get(key string) (string, error) {
	kvStore.mutex.RLock()
	defer kvStore.mutex.RUnlock()
	if value, ok := kvStore.internalMap[key]; ok {
		return value, nil
	} else {
		return "", fmt.Errorf("%s not found", key)
	}
}