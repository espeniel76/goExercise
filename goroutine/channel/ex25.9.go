package channel

import (
	"context"
	"fmt"
)

func Ex25_9() {
	wg.Add(1)
	ctx := context.WithValue(context.Background(), "number", 9)
	go square5(ctx)
	wg.Wait()
}

func square5(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square: %d", n*n)
	}
	wg.Done()
}
