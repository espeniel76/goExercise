/*
문제 설명
직사각형을 만드는 데 필요한 4개의 점 중 3개의 좌표가 주어질 때, 나머지 한 점의 좌표를 구하려고 합니다. 점 3개의 좌표가 들어있는 배열 v가 매개변수로 주어질 때, 직사각형을 만드는 데 필요한 나머지 한 점의 좌표를 return 하도록 solution 함수를 완성해주세요. 단, 직사각형의 각 변은 x축, y축에 평행하며, 반드시 직사각형을 만들 수 있는 경우만 입력으로 주어집니다.

제한사항
v는 세 점의 좌표가 들어있는 2차원 배열입니다.
v의 각 원소는 점의 좌표를 나타내며, 좌표는 [x축 좌표, y축 좌표] 순으로 주어집니다.
좌표값은 1 이상 10억 이하의 자연수입니다.
직사각형을 만드는 데 필요한 나머지 한 점의 좌표를 [x축 좌표, y축 좌표] 순으로 담아 return 해주세요.
*/

package examtest

import "fmt"

func Ex1() {
	v1 := [][]int{
		{1, 4},
		{3, 4},
		{3, 10},
	}
	fmt.Println(solution(v1))

	v2 := [][]int{
		{1, 1},
		{2, 2},
		{1, 2},
	}
	fmt.Println(solution(v2))
}

func solution(v [][]int) []int {
	answer := []int{0, 1}
	uesX := make(map[int]int)
	uesY := make(map[int]int)

	// x, y 숫자는 반드시 2개가 사용된다.
	for _, v1 := range v {
		// fmt.Println(v1[0], v1[1])
		// for k2, v2 := range v1 {
		// 	fmt.Println("step2", k2, v2)
		// }
		uesX[v1[0]]++
		uesY[v1[1]]++
	}
	// fmt.Println(uesX)
	// fmt.Println(uesY)
	for k, v := range uesX {
		// fmt.Println(k, v)
		if v < 2 {
			answer[0] = k
		}
	}
	for k, v := range uesY {
		// fmt.Println(k, v)
		if v < 2 {
			answer[1] = k
		}
	}

	return answer
}
