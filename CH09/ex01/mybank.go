// Copyright © 2016 "Shun Yokota" All rights reserved

// Package mybank provides a concurrency-safe bank with one account.
package mybank

type withdrawMsg struct {
	amount int
	ok     bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdrawMsg)

//Deposit is write deposits to chan.
func Deposit(amount int) { deposits <- amount }

//Balance is read balance from chan.
func Balance() int { return <-balances }

//Withdraw is
func Withdraw(amount int) bool {
	withdraws <- withdrawMsg{amount, false}
	msg := <-withdraws
	//fmt.Println(msg.amount)
	return msg.ok
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case msg := <-withdraws:
			if balance < msg.amount {
				withdraws <- withdrawMsg{0, false}
				break
			}
			balance -= msg.amount
			withdraws <- withdrawMsg{balance, true}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
