package examples

import "fmt"

func VariablesMethods() {

	// {}를 이용한 초기화 방법
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}

	fmt.Println(slice1)
	fmt.Println(slice2)

	// make 를 이용한 초기화 방법
	var slice = make([]int, 3)
	fmt.Println(slice)

	// slice 요소접근
	fmt.Println(slice[2])

	// slice 순회
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
	// range 의 첫번째 인자는 index, 두번째가 value 이다
	for i, v := range slice1 {
		fmt.Println(i, v)
	}
}
