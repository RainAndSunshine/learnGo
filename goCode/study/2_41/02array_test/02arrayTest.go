package main

import "fmt"

func main() {

	//求数组[1,3,5,7,8]所有元素的和
	a1 := [...]int{1, 3, 5, 7, 8}
	sum1 := 0
	for _, v := range a1 {
		sum1 += v
	}
	fmt.Println(sum1)

	//找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为（0,3）和（1,2）
	isOk := false
	for i := 0; i < len(a1); i++ {
		isOk = false
		for t := i + 1; t < len(a1); t++ {
			if a1[i]+a1[t] == 8 {
				fmt.Printf("(%d,%d)\n", i, t)
				isOk = true
			}
			if isOk {
				break
			}
		}
	}

}
