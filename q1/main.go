package main

import "sync"

type User struct {
	ID      uint64
	Balance uint64
	Lock    sync.Mutex
}

func transfer(from *User, to *User, amount uint64) {
	if from.ID < to.ID {
		from.Lock.Lock()
		defer from.Lock.Unlock()
		to.Lock.Lock()
		defer to.Lock.Unlock()
	} else {
		to.Lock.Lock()
		defer to.Lock.Unlock()
		from.Lock.Lock()
		defer from.Lock.Unlock()
	}

	if from.Balance >= amount {
		from.Balance -= amount
		to.Balance += amount
	}
}

func main() {
	userA := User{
		ID: 1, Balance: 10e10,
	}
	userB := User{
		ID: 2, Balance: 10e10,
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10e10; i++ {
			transfer(&userA, &userB, 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10e10; i++ {
			transfer(&userB, &userA, 1)
		}
	}()
	wg.Wait()
}
