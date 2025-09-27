package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	logSafe := func(goroutineID int, msg string) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Printf("[Горутина %d]: %v\n", goroutineID, msg)
	}
	for i := 1; i <= 1; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			logSafe(id, "начало работы")
			time.Sleep(100 * time.Millisecond)
			logSafe(id, "выполнение")
			time.Sleep(100 * time.Millisecond)
			logSafe(id, "завершила работу")
		}(i)
	}
	wg.Wait()
	fmt.Printf("логи записаны")
}
