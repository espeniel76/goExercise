/*
주어진 숫자 중 3개의 수를 더했을 때 소수가 되는 경우의 개수를 구하려고 합니다. 숫자들이 들어있는 배열 nums가 매개변수로 주어질 때, nums에 있는 숫자들 중 서로 다른 3개를 골라 더했을 때 소수가 되는 경우의 개수를 return 하도록 solution 함수를 완성해주세요.

제한사항
nums에 들어있는 숫자의 개수는 3개 이상 50개 이하입니다.
nums의 각 원소는 1 이상 1,000 이하의 자연수이며, 중복된 숫자가 들어있지 않습니다.
*/

package level1

import "fmt"

func Ex06() {
	fmt.Println(solution6([]int{1, 2, 3, 4}))
	// i := 0
	// for i = 1; i <= 1000; i++ {
	// 	if IsPrimeNumber(i) > 0 {
	// 		// printf("%3d ", i); //i를 출력
	// 		fmt.Printf("%3d ", i)
	// 	}
	// }
	// fmt.Println("\n")
}

// 소수점 구하는 함수
func IsPrimeNumber(n int) int {
	i := 0
	last := n / 2
	if n <= 1 {
		return 0
	}
	for i = 2; i <= last; i++ {
		if (n % i) == 0 {
			return 0
		}
	}
	return 1
}

func solution6(nums []int) int {

	for key, val := range nums {

	}

	return answer
}
