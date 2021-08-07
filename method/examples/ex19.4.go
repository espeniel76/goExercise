package examples

import "fmt"

type account3 struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a1 *account3) withdrawPointer(amount int) {
	a1.balance = a1.balance - amount
}

// 값 타입 메서드
func (a2 account3) withdrawValue(amount int) {
	a2.balance = a2.balance - amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account3) withdrawReturnValue(amount int) account3 {
	a3.balance -= amount
	return a3
}

func Ex19_4() {
	var mainA *account3 = &account3{100, "Joe", "Park"}

	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account3 = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawValue(30)
	fmt.Println(mainB.balance)
}
