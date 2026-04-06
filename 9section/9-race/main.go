package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.balance < amount {
		fmt.Println("insufficient funds")
		return
	}
	b.balance -= amount
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance: %d", b.balance)
}

func (b *BankAccount) GetBalance() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.balance
}


func main() {
	// help you prevent race condition.
	// When used correctly you can ensure that only one goroutine can access a variable at a time.

	// counter := 0 // critical section

	// var wg sync.WaitGroup
	// var mu sync.Mutex

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		mu.Lock()
	// 		counter++
	// 		mu.Unlock()
	// 	}()
	// }

	// wg.Wait()
	// fmt.Printf("counter: %d\n", counter)

	var wg sync.WaitGroup
	var account = &BankAccount{
		balance: 100,
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Deposit(amount)
		}(i + 1)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Withdraw(amount)
		}(i + 1)
	}

	wg.Wait()
	fmt.Println(account)
}
