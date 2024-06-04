package main

import "fmt"

const (
	eat   int = 1
	sleep int = 2
	da    int = 4
)

// 111
// 左边的1表示吃饭		100
// 中间的1表示睡觉 	010
// 右边的1表示打豆豆	001
func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | da) //101
}
