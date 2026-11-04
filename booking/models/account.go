package models

import "errors"

type BankAccount struct{
	Owner string // Public field
	balance float64  // Private field: cannot be accessed outside package "account"
}

// NewBankAccount acts as a constructor. 
// It validates inputs before returning a successfully initialized pointer.
func NewBankAccount(owner string,intialDeposit float64) (*BankAccount,error){

	if intialDeposit <0 {
		return nil,errors.New("initial deposit cannot be negative")
	}
	return &BankAccount{
		Owner:owner,
		balance:intialDeposit,// Package code CAN access private fields
	},nil

}

func (account *BankAccount)GetBalance()float64{
	return account.balance
}

func (account *BankAccount)Deposit(deposit float64) (float64,error){
	if deposit <=0 {
		return 0.0, errors.New("Deposit cannot be negative or zero")
	}
	 account.balance += deposit
	 return account.balance,nil
}

func (account *BankAccount)Withdraw(amount float64) (float64,error){
	if amount <=0 {
		return 0.0,errors.New("Amount cannot be negative or zero")
	}
	if account.balance < amount{
		return 0.0,errors.New("Insufficient balance")
	}

	account.balance -= amount

	return account.balance,nil
}

