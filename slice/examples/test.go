package examples

import "fmt"

func Test() {

	// 배열
	var array = [2]int{}
	fmt.Println(array)

	// 값 할당
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	// 생성과 동시에 할당
	array = [2]int{100, 200}
	fmt.Println(array)

	// 슬라이스
	var slice = []int{}
	fmt.Println(slice)
	slice = append(slice, 1)
	fmt.Println(slice)

	// append 로 할당된 영역에 값을 할당 할 수 있음
	slice[0] = 2
	fmt.Println(slice)

	// 이렇게 하면 오류남
	// 신규할당은 append 로 해야함
	// slice[1] = 3
	// fmt.Println(slice)
}
