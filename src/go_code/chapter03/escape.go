package main

import "fmt"

func main()  {
	//声明一个数组，长度为4，数据类型为int,定义数组长度后面不需要写等于号
	var arr1 [4] int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4

	//直接怼数组进行赋值  需要写等于号，根据后面的值自动推断个数
	var arr2 = [...] int {6,2,3}
	//fmt.Println(arr2[1])
	//使用for range 来循环数组 for 下标,值 := range 数组 { }
	//for index,value := range arr2{
	//	fmt.Printf("数组的下标是%d,值是%d",index,value)
	//	fmt.Println()
	//}
	for i := 1;i < len(arr2);i++ {
		for j := 0;j<len(arr2) - i;j++ {
			if arr2[j] > arr2[j+1] {
				arr2[j],arr2[j+1] = arr2[j+1],arr2[j]
			}
		}
	}
	fmt.Println(arr2)
	//var n [10]int /* n 是一个长度为 10 的数组 */
	//var i,j int
	//
	///* 为数组 n 初始化元素 */
	//for i = 0; i < 10; i++ {
	//	n[i] = i + 100 /* 设置元素为 i + 100 */
	//}
	//
	///* 输出每个数组元素的值 */
	//for j = 0; j < 10; j++ {
	//	fmt.Printf("Element[%d] = %d\n", j, n[j] )
	//}
}