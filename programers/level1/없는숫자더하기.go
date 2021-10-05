/*
0부터 9까지의 숫자 중 일부가 들어있는 배열 numbers가 매개변수로 주어집니다.
numbers에서 찾을 수 없는 0부터 9까지의 숫자를 모두 찾아 더한 수를 return 하도록 solution 함수를 완성해주세요.
*/
package level1

import "fmt"

func Ex03() {
	numbers := []int{1, 2, 3, 4, 6, 7, 8, 0}
	fmt.Println(solution3(numbers))
}

func solution3(numbers []int) int {

	compares := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	nones := []int{}

	isExist := false
	for _, val1 := range compares {
		isExist = false
		for _, val2 := range numbers {
			if val1 == val2 {
				isExist = true
				break
			}
		}
		if !isExist {
			nones = append(nones, val1)
		}
	}
	sum := 0
	for _, val := range nones {
		sum += val
	}

	return sum
}
