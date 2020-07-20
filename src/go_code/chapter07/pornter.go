package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("变量a的内存地址是%p",&a)
	fmt.Println()
	var p1 *int
	//将变量a的内存地址赋值给指针
	p1 = &a
	fmt.Printf("指针p1的内存地址是%p",&p1)
	fmt.Println()
	fmt.Printf("指针p1的内存地址指向的值是%d",*p1)
	fmt.Println()
	var p2 **int
	p2 = &p1

	fmt.Printf("指针p2的内存地址是%p",&p2)
	fmt.Println()
	fmt.Printf("指针p2的内存地址指向的值的值是%d",**p2)
	fmt.Println()
	//数组指针，类型是指针，存储的是数组的地址
	 var arr1 = [4]int{1,2,3,4}
	 var p3 = &arr1
	 fmt.Printf("arr1的地址是%p\n",&arr1)
	 fmt.Println("p3的值指向的值是",*p3)
	 //简写方式，访问p3地址指向的地址的值，更改为100
	 p3[0] = 100
	 fmt.Println(arr1)

	 //指针数组，数组里面存储的是指针
	 a = 1
	 b := 2
	 c := 3
	 d := 4
	 // 数组的大小，存放int 类型指针的
	 arr2 := [4]*int{&a,&b,&c,&d}
	 fmt.Println(*arr2[0])
	 *arr2[0] = 7
	 fmt.Println(a)
}
