package main

import "fmt"

//全局变量
var(
	n4 = 111
	n5 = "jjj"
)

func main()  {
	var i int
	i = 10
	fmt.Println("i=", i)

	//多个变量
	n1, n2, n3 := 100, "haha", 99
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	fmt.Println("n4=", n4, "n5=", n5)
}