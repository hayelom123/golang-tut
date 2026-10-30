package main

import ("fmt" 
"sync")


func printNumbers(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=1; i<3; i++{
		fmt.Println("Number: ",i)
	}
}


func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go printNumbers(&wg)

	wg.Wait()
	fmt.Println("Finished!")

}