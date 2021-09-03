package examples

import "fmt"

func Ex14_1() {
	var a int = 500
	fmt.Printf("a의 값: %d\n", a) // 500
	p := &a
	fmt.Printf("p의 값: %p\n", p)            // 0xc0000aa068
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) // 500
	*p = 100
	fmt.Printf("a의 값: %d\n", a) // 100
}
