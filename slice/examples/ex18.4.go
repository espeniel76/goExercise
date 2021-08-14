package examples

import "fmt"

// 배열은 배열 자체가 복사된다.
// 고로, 복사본의 2 index 값이 200이 할당된다.
func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeArrayByRef(array2 *[5]int) {
	array2[2] = 200
}

// slice 는 배열의 주소값이 복사된다.
// 고로, 2 index 값을 200으로 하면, 기본값이 변경된다.
func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func Ex18_4() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	fmt.Println("array:", array)

	changeArrayByRef(&array)
	fmt.Println("array:", array)

	changeSlice(slice)
	fmt.Println("slice:", slice)
}
