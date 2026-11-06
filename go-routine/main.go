package main

import (
	"fmt"
	// "time"
	"sync"
)



var wg sync.WaitGroup


func display(task string){
 fmt.Println(task)
 wg.Done()
}
func worker() {
	defer wg.Done()
    for i := 1; i <= 5; i++ {
		wg.Add(1)
        go display(fmt.Sprintf("Processing task: %v", i))
    }
}

func main(){
     wg.Add(1)
	go  worker()
    
	// time.Sleep(time.Second *6)
     wg.Wait()
	fmt.Println("Done")

}