package main

import (
	"fmt"
	// "reflect"
	// "unsafe"
)
func main()  {
	// num1 := 10
	// fmt.Printf("num1的数据类型是 %T", num1)
	// fmt.Println(reflect.TypeOf(num1))
	// fmt.Println(unsafe.Sizeof(num1))
	// a1 := 'a'
	// fmt.Printf("a1的值是: %d", a1)
	// a1 := 111
	// a2 := 11.11
	// a3 := true
	var a3 byte = 'h'
	//数据类型转换
	// var a2 float32 = float32(a1)
	// fmt.Printf("type %T, a2=%v", a2,a2)
	//数值转字符串
	var n3 string
	// a3 = fmt.Sprintf("%d", a1)
	// a3 = fmt.Sprintf("%f", a2)
	n3 = fmt.Sprintf("%c", a3)
	fmt.Printf("a3 type %T a3=%q\n", n3, n3)
}