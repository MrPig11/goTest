package main

import "fmt"

//定义一个全局的变量


func main() {
	fmt.Println("打印第一句话")
	// defer用于延迟执行
	defer fmt.Println("打印第二句话")
	fmt.Println("打印第三句话")
	// defer 存储采用栈的操作  先进后出
	defer fmt.Println("打印第四句话")
	i:=fun3()
	fmt.Println(i)
}

func fun3() int {
	fmt.Println("打印")
	defer fmt.Println("第二次打印")
	// 此时return 终止函数执行了  所以先打印上面的defer
	return 0
}

