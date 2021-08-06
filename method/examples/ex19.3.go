package examples

import "fmt"

type account2 struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a1 *account2) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account2) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account2) withdrawReturnValue(amount int) account2 {
	a3.balance -= amount
	return a3
}

func Ex19_3() {
	var mainA *account2 = &account2{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account2 = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}
