package main

import (
	"sync"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.store[name]
}
func (i *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}
func (i *InMemoryPlayerStore) close() {}
