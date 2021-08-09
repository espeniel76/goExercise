package examples

import "fmt"

func Ex18_1() {
	// slice 는 배열의 길이를 특정하지 않고 사용할 수 있따.
	var slice []int

	// slice 길이가 0인지 확인
	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10 // 하지만, panic 발생
	fmt.Println(slice)

	/* 즉, slice 를 하려면
	초기화를 해 주어야 한다. */

}
