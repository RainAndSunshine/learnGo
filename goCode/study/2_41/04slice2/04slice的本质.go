package main

import "fmt"

//切片的本质
//切片就是一个框，框柱了一块连续的内存
//属于引用类型，真正的数据都是保存在底层数组里的
//切片不存值

func main() {
	//make()函数创造切片
	//make(数据类型，长度，容量)
	//不写容量则会默认容量与长度一致
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s2, len(s2), cap(s2))

	//切片是不能直接比较的，唯一能用 == 的是和nil比较
	//一个nil值的切片的长度和容量都是0，但我们不能说一个长度和容量都是0的切片一定是nil
	//一个nil值的切片是没有底层数组的
	//所以要判断一个切片s是否为空，要用len(s)==0来判断
	var s11 []int         //len(s11)==0;cap(s11)==0;s11==nil
	s22 := []int{}        //len(s22)==0;cap(s22)==0;s22!=nil
	s33 := make([]int, 0) //len(s33)==0;cap(s33)==0;s33!=nil
	fmt.Println(s11, s22, s33)

	//切片的赋值拷贝
	//拷贝前后的两个切片共享底层数组，对一个切片的修改会影响另一个切片的内容
	s44 := make([]int, 5) //[0,0,0,0,0]
	s55 := s44
	s44[4] = 999
	fmt.Println(s44)
	fmt.Println(s55)

	//切片的遍历
	//1.索引遍历
	for i := 0; i < len(s44); i++ {
		fmt.Println(s44[i])
	}
	//2.for range遍历
	for i, v := range s44 {
		fmt.Println(i, v)
	}
}
