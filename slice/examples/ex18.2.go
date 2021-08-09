package examples

import "fmt"

func Ex18_2() {
	var slice = []int{1, 2, 3}

	// slice 에 요소추가
	slice = append(slice, 4)
	slice = append(slice, 6)
	fmt.Println(slice)

	// 당연 복사를 통한 추가 생성도 가능
	slice2 := append(slice, 7)
	fmt.Println(slice2)

	var slice3 []int
	for i := 0; i <= 10; i++ {
		slice3 = append(slice3, i)
	}
	slice3 = append(slice3, 11, 12, 13, 14, 15, 16)

	fmt.Println(slice3)

}
