package channel

import (
	"fmt"
	"sync"
	"time"
)

func square4(wg *sync.WaitGroup, ch chan int) {

	tick := time.Tick(time.Second)
	terminate := time.After(time.Second * 100)

	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: [%d] %d\n", n, n*n)
			time.Sleep(time.Microsecond)

		case <-tick:
			// fmt.Println("Tick")

		case <-terminate:
			fmt.Println("Terminate")
			wg.Done()
			return
		}
	}
}

func Ex25_6() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square4(&wg, ch)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	wg.Wait()
}
