package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker processes jobs from the jobs channel until it is empty or context is canceled
func worker(id int, ctx context.Context, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[Worker %d] Aborting: context canceled\n", id)
			return

		case job, open := <-jobs:
			if !open {
				// The jobs channel was closed normally by main()
				fmt.Printf("[Worker %d] Finished all assigned work.\n", id)
				return
			}

			// Simulate processing the job (takes 300ms)
			fmt.Printf("[Worker %d] Started job %d\n", id, job)
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("[Worker %d] Completed job %d\n", id, job)
		}
	}
}

func main() {
	// 1. Setup global control (cancellation after 1 second)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 2. Setup job queue and synchronization
	numJobs := 10
	numWorkers := 3
	jobsChan := make(chan int, numJobs)
	var wg sync.WaitGroup

	// 3. Spin up 3 workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, ctx, jobsChan, &wg)
	}

	// 4. Fill the queue with 10 jobs
	for j := 1; j <= numJobs; j++ {
		jobsChan <- j
	}
	close(jobsChan) // Signal that no more jobs will ever be added

	// 5. Wait for workers to finish or time out
	wg.Wait()
	fmt.Println("Main thread finished execution flow.")
}
