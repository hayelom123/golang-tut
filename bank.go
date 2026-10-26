
package main

import (
    "errors"
    "fmt"
)

type BankAccount struct {
    Owner   string
    Balance float64
}

func (a *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("deposit must be positive")
    }

    a.Balance += amount
    return nil
}

func (a *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("withdrawal must be positive")
    }

    if amount > a.Balance {
        return errors.New("insufficient funds")
    }

    a.Balance -= amount
    return nil
}

func main() {
    account := BankAccount{
        Owner:   "Alex",
        Balance: 5000,
    }

    if err := account.Deposit(2000); err != nil {
        fmt.Println("Deposit error:", err)
        return
    }

    if err := account.Withdraw(1500); err != nil {
        fmt.Println("Withdrawal error:", err)
        return
    }

    fmt.Println("Final balance:", account.Balance)
}
