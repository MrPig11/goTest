package main

import "fmt"

func main()  {
	//score := 0
	//fmt.Scanln(&score)
	//if score > 60 {
	//	fmt.Println("您及格了")
	//} else {
	//	fmt.Println("您不及格")
	//}
	//if num := 1; num == 1 {
	//	fmt.Println("是真的")
	//} else {
	//	fmt.Println("num不等于1")
	//}
	//fmt.Println(num)
	//num := 1
	//switch num {
	//case 1,2,3,4:
	//	fmt.Println("1111")
	//	// 穿透case
	//	fallthrough
	//case 5:
	//	fmt.Println("5555555")
	//
	//}
	score := 60
	switch {
	case score < 60:
		fmt.Println("不及格")
	case score >= 60:
		fmt.Println("及格")

	}
}