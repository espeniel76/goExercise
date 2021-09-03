package examples

import "fmt"

type Data2 struct {
	value int
	data  [200]int
}

func changeData2(arg *Data2) {
	arg.value = 900
	arg.data[100] = 999
}

func Ex14_4() {
	var data Data2

	changeData2(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
