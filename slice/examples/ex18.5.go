package examples

import "fmt"

func Ex18_5() {
	fmt.Println("##### ex18.5 #####")
	// slice 는 기본적으로 같은 배열을 바라보는 주소값이다.
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100
	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)
	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
