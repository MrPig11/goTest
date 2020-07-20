package main

import (
	"fmt"
)

func main()  {
	// var i int = 10
	// fmt.Println("i的地址是", &i)
	// var p *int = &i
	// fmt.Println("p的地址是",p)
	// fmt.Println("p的值为",*p)
	var name string
	var age int
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("是否通过考试")
	fmt.Scanln(&isPass)
	fmt.Printf("\n 姓名是%v\n年龄是%v\n是否通过考试%v", name,age,isPass)
}