package examples

import "fmt"

func Ex18_6() {
	fmt.Println("##### ex18.6 #####")
	// 같은 주소값 이지만, 공간이 부족하면 더 넓은 공간을 새로 만들어 버린다.
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

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
