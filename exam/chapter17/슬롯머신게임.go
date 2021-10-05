/*
간단한 슬롯머신 게임을 만들어보겠습니다. 가진 돈은 1000원으로 시작합니다.
1~5 사이의 값을 입력받습니다. 그런 뒤 1~5사이 랜덤한 값을 선택합니다.
만약 입력한 값과 랜덤한 값이 같으면 가진 돈에 500원을 추가하고 축하한다는 메시지와 잔액을 표시합니다.
다를 경우 가진 돈에서 100원을 빼고 아쉽다는 메시지와 잔액을 표시합니다. 다시 1~5사이의 값을 입력받는 것을 반복 하다가
가진 돈이 0원 이하가 되거나 5000원 이상이 되면 게임을 종료합니다.
*/
package chapter17

import (
	"fmt"
	"math/rand"
	"time"
)

const MSG_BETWEEN = "1~5사이의 값을 입력하세요."

func Exam02() {
	remainCash := 1000 // 잔액

	for {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(5)
		fmt.Println("1~5 사이의 값을 입력 받습니다.", r)
		n, err := InputItValud()
		if err != nil {
			fmt.Println("숫자만 입력 하세요.")
		} else {
			if n < 1 {
				fmt.Println(MSG_BETWEEN)
			} else if n > 5 {
				fmt.Println(MSG_BETWEEN)
			} else {
				if r == n {
					remainCash += 500
					fmt.Println("축하합니다. 잔액:", remainCash)
				} else {
					remainCash -= 100
					fmt.Println("아쉽습니다. 잔액:", remainCash)
				}

				if remainCash >= 5000 {
					fmt.Println("잔액이 5000을 넘어 게임을 종료 합니다:", remainCash)
					break
				}
				if remainCash <= 0 {
					fmt.Println("잔액이 0이 되어 게임을 종료 합니다:", remainCash)
					break
				}
			}

		}
	}
}
