package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	// Tell the WaitGroup this worker is finished when the function returns
	defer wg.Done() 
	
	fmt.Printf("Worker %d starting\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter for each background task
		go worker(i, &wg)
	}
 
	wg.Wait() // Blocks execution here until the counter hits 0
	fmt.Println("All workers finished!")
}
