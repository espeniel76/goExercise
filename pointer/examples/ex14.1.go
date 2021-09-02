package examples

import "fmt"

func Ex14_1() {
	var a int = 500
	fmt.Printf("a의 값: %d\n", a)
	p := &a
	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)
	*p = 100
	fmt.Printf("a의 값: %d\n", a)
}
