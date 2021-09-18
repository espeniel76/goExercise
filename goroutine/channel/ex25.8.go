package channel

import (
	"context"
	"fmt"
	"time"
)

// var wg sync.WaitGroup

func Ex25_8() {
	wg.Add(1)
	// 컨텍스트 생성
	ctx, cancel := context.WithCancel(context.Background())
	go printEverySecond(ctx)
	time.Sleep((5 * time.Second))
	cancel()

	wg.Wait()
}
func Ex25_81() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go printEverySecond(ctx)
	time.Sleep((5 * time.Second))
	cancel()

	wg.Wait()
}

func printEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context")
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
