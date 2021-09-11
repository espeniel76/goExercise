package exercise

import (
	"fmt"
	"time"
)

func printHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func Ex24_1() {
	go printHangul()
	go printNumbers()

	time.Sleep(3 * time.Second)
}
