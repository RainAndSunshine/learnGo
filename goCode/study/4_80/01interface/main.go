package main

import "fmt"

//接口 interface
//接口是一种类型，是一种特殊的类型，它规定了变量有哪些方法
//在编程中会遇到以下场景：
//我不关心一个变量是什么类型，我只关心能调用它的什么方法

// 引出接口的实例
// 定义一个接口speaker类型
type speaker interface {
	speak()
}

type cat struct{}
type dog struct{}
type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("啊啊啊")
}

func da1(x dog) {
	//接受一个参数，传进来什么，我就打什么
	x.speak() //挨打了就要叫
}

func da2(x speaker) {
	x.speak()
}

func main() {
	//只有当有两个或者两个以上的具体类型必须以相同的方式进行处理时才需要定义空接口
	//不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗

	var (
		cat    cat
		dog    dog
		person person
	)
	//使用方法da1
	//da1(cat)	//因为方法限制了参数只能是狗，所以传不进去
	//da1(dog)
	//da1(person)

	//使用接口da2
	//使用接口类型，拥有speak方法的都可以进入接口da2
	da2(cat)
	da2(dog)
	da2(person)
}
