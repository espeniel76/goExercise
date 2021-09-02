package examples

import "fmt"

func Test01() {
	var a int
	var p *int
	a = 20
	fmt.Println(a)
	p = &a
	fmt.Println(p)
	*p = 100
	fmt.Println(p)
	fmt.Println(a)
}
