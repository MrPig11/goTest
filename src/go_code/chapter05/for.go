package main

import (
	"fmt"
	"math"
)

func main() {
	//for i := 1; i <= 5; i++ {
	//	fmt.Println("循环输出")
	//}
	//a := 1
	//for a <= 5 {
	//	fmt.Println("我是变量a的循环")
	//	a++
	//}
	//b := 1
	//不写条件相当于while(true)
	//for  {
	//	fmt.Printf("变量b的数据为%d",b)
	//	b++
	//}
	//for i := 1; i <= 5; i++{
	//	for j := 1; j <= 5; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	// 99乘法表
	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d * %d = %d\t",i,j,i*j)
	//	}
	//	fmt.Println()
	//}

	//水仙花数，每个数字上的立方和相加等于数字本身
	for i := 100; i < 1000; i++ {
		//百位
		a := i / 100

		//10位
		b := i / 10 % 10
		//各位
		c := i % 100 % 10
		if math.Pow(float64(a),3) + math.Pow(float64(b),3) + math.Pow(float64(c),3) == float64(i) {
			fmt.Println(i)
		}
		//fmt.Println()
	}
}
