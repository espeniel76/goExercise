package channel

import (
	"fmt"
	"sync"
	"time"
)

func square3(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			fmt.Println("Quit!")
			return
		}
	}
}

func Ex25_5() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square3(&wg, ch, quit)

	for i := 0; i < 5; i++ {
		ch <- i * 2
	}
	quit <- true
	wg.Wait()
}
