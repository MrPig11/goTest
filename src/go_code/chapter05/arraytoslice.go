package main

import "fmt"

func main() {
	arr1 := [10] int {0,1,2,3,4,5,6,7,8,9}
	s1 := arr1[0:5] //从0下标到5下标-1，包前不包尾
	s2 := arr1[6:]//从6下标到最后 len
	s3 := arr1[:]//全部
	s4 := arr1[3:8]//从3到8下标-1
	s5 := arr1[:8]//从0到8下标-1 左指针没动，容量为10
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println("----------切边本身不存储数据，存储的是数组的内存地址------------")
	fmt.Println("----------更改数组的值，切片的值也会更改，更改切片的值，数组的值也会更改------------")
	arr1[0] = 100
	fmt.Println(arr1)
	fmt.Println(s1)
	s2[2] = 90
	fmt.Println(arr1)
	fmt.Println(s2)
	fmt.Println("---------将数组转为切片时,切片的容量为原来的数组的容量减数组左指针经过的数--------------")
	fmt.Printf("s1切片的长度为%d,容量为%d",len(s1),cap(s1))
	fmt.Println()
	fmt.Printf("s2切片的长度为%d,容量为%d",len(s2),cap(s2))
	fmt.Println()
	fmt.Printf("s3切片的长度为%d,容量为%d",len(s3),cap(s3))
	fmt.Println()
	fmt.Printf("s4切片的长度为%d,容量为%d",len(s4),cap(s4))
	fmt.Println()
	fmt.Printf("s4切片的长度为%d,容量为%d",len(s5),cap(s5))
}
