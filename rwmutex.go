package main

import ("fmt"
		"sync"
		)

type BankAccount struct{
	balance int
    mu      sync.RWMutex
}

func (a *BankAccount) GetBalance()int{
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.balance
}

func (a *BankAccount) Deposit(amount int){
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance+=amount
}


func main(){

	account := &BankAccount{
					balance:1000,
				}

	var wg sync.WaitGroup
	wg.Add(2)

	for i:=1; i<=2;i++{
		go func(id int){
			defer wg.Done()

			fmt.Println(
                "Reader", id,
                "Balance:", account.GetBalance(),
            )
		}(i)
	}
	wg.Wait()

	account.Deposit(500)

	 fmt.Println("Final balance:", account.GetBalance())

}