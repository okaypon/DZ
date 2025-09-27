package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	pageViews int
	mu        sync.Mutex
)

func registerPageView() {
	mu.Lock()
	pageViews++
	mu.Unlock()
}

func simulateUser(polzovatel int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		registerPageView()
		fmt.Printf("пользлватель %d просмотрел страницу\n", polzovatel)
		time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go simulateUser(i, &wg)
	}
	wg.Wait()
	fmt.Printf("просмотров страницы: %d\n", pageViews)
}
