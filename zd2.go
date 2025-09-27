package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func main() {
	cache := Cache{data: make(map[string]string)}
	cache.mu.Lock()
	cache.data["ffff"] = "aga"
	cache.mu.Unlock()
	cache.mu.RLock()
	polzovatel := cache.data["ffff"]
	cache.mu.RUnlock()
	fmt.Println(polzovatel)
}
