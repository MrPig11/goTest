package main

import "fmt"

func main()  {
	//转义字符
	fmt.Println("aaa\tbbb")
	//换行符
	fmt.Println("aaa\nbbb")
	//将符号转义成字符串
	fmt.Println("张三说：\"早上好\"")
	//回车符， 先把回车符之前的内容输出，再用后面的字符串替换
	fmt.Println("天龙八部雪山飞狐\r刘备")
	/*
	我是块注释
	*/
	//小练习
	fmt.Println("姓名\t年龄\t籍贯\t住址\n张三\t12\t河北\t山东")
}