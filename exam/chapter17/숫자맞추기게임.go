/*
1. 0~99 사이의 랜덤한 숫자 하나를 정합니다.
2. 사용자 입력을 받습니다.
3. 입력값과 랜덤값을 비교합니다. 만약 사용자 입력 숫자가 더 크다면
"입력하신 숫자가 더 큽니다" 를 출력하고 작으몀
"입력하신 숫자가 더 작습니다." 를 출력합니다.
다시. 사용자 입력을 받아서 반복합니다.
4. 만약 숫자가 맞으면 "축하합니다. 숫자를 맞추셨습니다. 시도횟수: 13번"
같이 메시지를 출력합니다.
5. 프로그램을 종료합니다.
*/

package chapter17

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputItValud() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func Exam01() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100000)
	cnt := 1
	for {
		fmt.Printf("숫자값을 입력하세요:")
		n, err := InputItValud()
		if err != nil {
			fmt.Println("숫자만 입력 하세요")
		} else {
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췃씁니다. 축하합니다. 시도한 횟수:", cnt)
				break
			}
		}
		cnt++
	}

}
