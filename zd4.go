package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var golosa int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(idgolosa int) {
			defer wg.Done()
			time.Sleep(time.Duration(idgolosa) * 100 * time.Millisecond)

			mu.Lock()
			golosa++
			mu.Unlock()

			fmt.Printf("пользователь %d проголосовал\n", idgolosa)
		}(i)
	}

	wg.Wait()
	fmt.Printf("подсчёт голосов: %d\n", golosa)
}
