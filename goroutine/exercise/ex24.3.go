package exercise

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func Ex24_3() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				depositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func depositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}