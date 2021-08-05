package examples

import "fmt"

// int 와 같은 내장 타입들도
// 별칭 myInt 라는 별칭 타입을 활용하여
type myInt int

// 메서드를 가질 수 있습니다.
func (a myInt) add(b int) int {
	return int(a) + b
}

func Ex19_2() {
	// myInt 타입 변수 a 인스턴스 생성
	var a myInt = 10

	// myInt 타입 변수 a 인스턴스에서 add 메서드 호출
	fmt.Println(a.add(30))

	var b int = 20

	// b 를 myInt 타입으로 변환 수 add 메서드 호출
	fmt.Println(myInt(b).add(50))

	// 이처럼 모든 리시버 타입이 별칭 타입으로 변환하여 메서드를 선언할 수 있습니다.
}
