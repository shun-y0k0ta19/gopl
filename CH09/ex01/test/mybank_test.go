// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"golang_training/CH09/ex01"
	//".."//use this import path, not in GOPATH
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		mybank.Deposit(200)
		fmt.Println("=", mybank.Balance())
		mybank.Withdraw(400)
		fmt.Println("=", mybank.Balance())
		mybank.Deposit(400)
		fmt.Println("=", mybank.Balance())
		mybank.Withdraw(400)
		fmt.Println("=", mybank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		mybank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := mybank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
