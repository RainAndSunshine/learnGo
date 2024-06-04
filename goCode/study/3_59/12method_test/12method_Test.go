package _2method_test

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看\新增学生\删除学生
*/

// 定义一个学生结构体
type student struct {
	id   int64
	name string
}

// 学生系统管理员
type manager struct {
	name        string
	allStudents map[int64]student
}

// 学生构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

// 菜单方法
func (m manager) menu() {
	fmt.Println()
	fmt.Println("欢迎使用学生信息管理系统")
	fmt.Print(`
-----------------------------
	1.查看学生信息
	2.增加学生信息
	3.删除学生信息
	4.退出管理系统
-----------------------------
	`)
}

// 选择方法
func (m manager) choice() {
	var c int
	var time int
	fmt.Print("请输入你的选择：")
	fmt.Scanln(&c)
	switch c {
	case 1:
		m.showAllStudents()
	case 2:
		m.addStudent()
	case 3:
		m.deleteStudent()
	case 4:
		os.Exit(1)
	default:
		fmt.Println("你的输入有误，请重试...")
		time++
		if time >= 3 {
			fmt.Println("错误次数过多，已自动退出...")
			os.Exit(1)
		}
	}
}

// 查看学生信息方法
func (m manager) showAllStudents() {
	fmt.Println()
	if len(m.allStudents) == 0 {
		fmt.Println("暂无学生信息...")
		return
	}
	for k, v := range m.allStudents {
		fmt.Printf("学号：%-8d\t姓名：%-8s\n", k, v.name)
	}
}

// 增加学生信息方法
func (m manager) addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入要增加的学生学号：")
	fmt.Scanln(&id)
	if _, ok := m.allStudents[id]; ok {
		fmt.Println("该学生信息已存在!!!")
		return
	}
	fmt.Print("请输入要增加的学生姓名：")
	fmt.Scanln(&name)
	tempS := newStudent(id, name)
	m.allStudents[id] = *tempS
	fmt.Println("添加成功！")
}

// 删除学生信息方法
func (m manager) deleteStudent() {
	var id int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scanln(&id)
	if _, ok := m.allStudents[id]; ok {
		delete(m.allStudents, id)
		fmt.Println("删除成功！")
	} else {
		fmt.Println("该学生不存在！！！")
		return
	}
}

func main() {
	var m manager
	fmt.Print("请输入管理员名称：")
	fmt.Scanln(&m.name)
	fmt.Printf("欢迎%s\n", m.name)
	m.allStudents = make(map[int64]student, 50)
	for {
		m.menu()
		m.choice()
	}
}
