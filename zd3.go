package main

import (
	"fmt"
	"sync"
)

func main() {

	queue := []string{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		queue = append(queue, "задача 1")
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		queue = append(queue, "задача 2")
		mu.Unlock()
	}()
	wg.Wait()
	mu.Lock()
	for _, ts := range queue {
		fmt.Println("обработана:", ts)
	}
	mu.Unlock()
	queue = []string{}
	fmt.Println("очередь пуста")
}
