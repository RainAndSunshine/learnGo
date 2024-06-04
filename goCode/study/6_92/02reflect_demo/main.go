package main

import (
	"fmt"
	"reflect"
)

//反射是指在程序运行期间对程序本身进行访问和修改的能力
//反射是一种原理，不需要自己去写
//利用反射可以将传入的空接口类型，取出里面的动态类型和动态值，从而进行访问和修改

//reflect包
//任何接口都是由 一个具体类型 和 具体类型的值 两部分组成
//任何接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成
//reflect包提供了reflect.Type和reflect.Value两个函数来获取任意对象的Value和Type

type Cat struct {
}

// 在反射中，类型还分为两种：类型Type 和 种类Kind
// 因为在Go语言中，我们可以用type关键字构造很多自定义类型，而种类Kind就是底层的类型

// 反射类型
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
	fmt.Printf("type name:%v, type kind:%v\n", v.Name(), v.Kind())
}

// 反射值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 通过反射设置变量的值
// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值
// 而反射中使用专有的Elem()方法来获取指针对应的值

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a float32 = 3.14
	reflectType(a) //type:float32
	var b int64 = 100
	reflectType(b) //type:int64

	var c = Cat{}
	reflectType(c)

	reflectValue(a)

	//reflectSetValue1(&b) //函数里没有Elem() 传指针也不行
	reflectSetValue2(&b)
	fmt.Println(b)
}
