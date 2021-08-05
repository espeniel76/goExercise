package examples

import "fmt"

type account struct {
	balance int
}

// 함수
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 구조체에 메서드 붙임
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func Ex19_1() {

	// 구조체 초기화
	a := &account{100}

	// 함수를 통한 차감
	withdrawFunc(a, 30)

	// 구조체의 메서드를 통한 차감
	a.withdrawMethod(30)
	fmt.Printf("%d \n", a.balance)
}
