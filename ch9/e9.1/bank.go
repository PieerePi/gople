// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type WITHDRAW struct {
	amount int
	result chan bool
}

var deposits = make(chan int)       // send amount to deposit
var balances = make(chan int)       // receive balance
var withdraws = make(chan WITHDRAW) // send amount to withdraw and receive result(true/false)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	result := make(chan bool)
	withdraws <- WITHDRAW{amount, result}
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		// If balances is an unbuffered channel and there is someone
		// blocking to read, or if balances is a buffered channel and
		// there is a vacancy, balance will be sent to balances,
		// otherwise this branch will not be selected.
		case balances <- balance:
		case withdraw := <-withdraws:
			if withdraw.amount <= balance {
				balance -= withdraw.amount
				withdraw.result <- true
			} else {
				withdraw.result <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
