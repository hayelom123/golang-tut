
package main

import (
    "fmt"
    "sync"
)

func main() {
    counter := 0

    var wg sync.WaitGroup
	var mu sync.Mutex

    wg.Add(2)

    for i := 0; i < 2; i++ {
        go func() {
            defer wg.Done()

            for j := 0; j < 100; j++ {
				mu.Lock()
                counter++
				// fmt.Println("func=================",i+1)
				mu.Unlock()
            }
        }()
    }

    wg.Wait()

    fmt.Println("Counter:", counter)
}
