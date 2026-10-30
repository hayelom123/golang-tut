package main

import ("fmt" 
"sync")

func sayHello() {
    fmt.Println("Hello!")
}

func sayGoodbye() {
    fmt.Println("Goodbye!")
}

func printNumbers(wg *sync.WaitGroup){

	defer wg.Done()

	for i:=1; i<=3; i++{
		fmt.Println("Number: ",i)
	}
}

func worker(name string,wg *sync.WaitGroup){
   defer wg.Done()
	fmt.Println("Hello from", name)
}

func main(){


//    go sayHello()
	// sayGoodbye()

	  var wg sync.WaitGroup

    wg.Add(2)

    go worker("Worker 1", &wg)
    go worker("Worker 2", &wg)

    wg.Wait()

    fmt.Println("All tasks completed!")

	// var wg sync.WaitGroup

	// wg.Add(1)

	// go printNumbers(&wg)

	// wg.Wait()

	// fmt.Println("Finished!")

}