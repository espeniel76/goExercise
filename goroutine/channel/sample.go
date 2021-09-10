package channel

import "fmt"

func Sample01() {
	var messages chan string = make(chan string)
	fmt.Println(messages)
}
