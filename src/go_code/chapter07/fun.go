package main

import "fmt"
//定义一个全局的变量
var l  = 1000

func main() {
	sum := fun1(10)
	fmt.Println(sum)

	sum2 := fun2(10,20,39)
	fmt.Println(sum2)
	fmt.Println(l)
}

func fun1(n int)int  {
	sum := 0
	for i := 1; i<= n;i++ {
		sum += i
	}
	return sum
}

func fun2(n ... int)(int) {
	sum := 0
	for _,v := range n{
		sum += v
	}
	return sum
}
