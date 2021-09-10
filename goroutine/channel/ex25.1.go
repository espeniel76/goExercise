package channel

import (
	"fmt"
	"sync"
	"time"
)

func Ex25_01() {

	// int 형 채널을 ch 이름으로 만들어라
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}
