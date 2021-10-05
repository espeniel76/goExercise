/*
두 문자열의 앞뒤를 겹쳐서 만들 수 있는 문자열 중, 더 짧은 문자열을 구하려합니다.
예를 들어 "xyZA"와 "ABCxy"가 주어졌을 때, 두 문자열을 겹치는 방법은 다음과 같습니다.
방법1. "xyZA" 뒤에 "ABCxy" 겹치기
방법2. "ABCxy" 뒤에 "xyZA" 겹치기
두 문자열 s1과 s2가 주어질 때, s1과 s2를 겹쳐서 만들 수 있는 문자열 중, 가장 짧은 문자열을 return 하도록 solution 함수를 완성해주세요.

제한사항
- s1과 s2의 길이는 1 이상 100 이하입니다.
- s1과 s2는 알파벳 대문자와 소문자로만 이루어져 있습니다.
- 문자열을 겹칠 때에는 대소문자를 구분합니다. 즉, "a"와 "A"를 겹칠 수는 없습니다.
- 가장 짧은 문자열이 여러 개라면 그 중 사전 순으로 앞서는 문자열을 return 해주세요.
*/
package examtest

import (
	"fmt"
)

func Ex2() {
	str1 := "xyZA"
	str2 := "ABCxy"
	result := solution2(str1, str2)
	fmt.Println(result)

	// str1 = "AxA"
	// str2 = "AyA"
	// result = solution2(str1, str2)
	// fmt.Println(result)
}

func solution2(s1 string, s2 string) string {
	result := ""
	s1StartNum := s1[0]
	s2Startnum := s2[0]
	firstString := s1
	secondString := s2

	// 선을 정한다.
	if s1StartNum > s2Startnum {
		firstString = s2
		secondString = s1
	} else {
		firstString = s1
		secondString = s2
	}

	for _, v1 := range firstString {
		isExist := false
		fmt.Println("string 1", string(v1))
		for _, v2 := range secondString {
			fmt.Println("string2", string(v2))

			if v1 == v2 {
				isExist = true
				result += string(v1)
				fmt.Println(result)
				break
			}
		}

		if !isExist {
			result += string(v1)
		}
	}

	// B조회 (B 남은부분 채운다)
	for _, v1 := range secondString {
		isExist := false
		for _, v2 := range result {
			if v1 == v2 {
				isExist = true
				break
			}
		}
		if !isExist {
			result += string(v1)
		}
	}

	return result
}
