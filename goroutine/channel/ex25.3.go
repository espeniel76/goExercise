package channel

import (
	"fmt"
	"sync"
	"time"
)

func square2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { //데이터를 모두 처리하고 난 다음에 채널이 딛힌 상태이면, for 문을 종료해서 프로그램이 정상 종료될 수 있다.
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func Ex25_03() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square2(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 데이터를 모두 넣고 채널이 더는 필요없기 때문에 close(ch)를 호출해 닫아줍니다.
	wg.Wait()
}
