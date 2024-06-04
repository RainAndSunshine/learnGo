package main

import "fmt"

//方法

type dog struct {
	name string
	age  int
}

func newDog(name string) *dog {
	return &dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
// 接受者是狗，函数的形参
func (d dog) barking() {
	fmt.Printf("%s：汪汪汪~\n", d.name)
}

// 值接受者和指针接受者的区别
//一个接受者的方法尽量统一用值接受或者指针接受

// 值接受者传的是拷贝
func (d dog) birth1() {
	d.age++
}

// 指针接受者传的是地址
func (d *dog) birth2() {
	d.age++
}

func main() {
	d1 := newDog("周琳")
	d1.barking()
	d1.age = 1
	fmt.Println(d1.age)
	d1.birth1()
	fmt.Println("经过值接受者的方法后:", d1.age)
	d1.birth2()
	fmt.Println("经过指针接受者的方法后:", d1.age)
}
