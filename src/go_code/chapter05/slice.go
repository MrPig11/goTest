package main

import "fmt"

func main() {
	// 切片(slice)，可变长的数组，切片是地址传递，数组是值传递
	//s1 := [] int {1,2,3}
	//s2 := s1
	//s2[0] = 8
	//fmt.Println(s1)
	//fmt.Println(s2)
	s1 := make([]int,3,3)
	fmt.Printf("s1的内存地址为%p",s1)
	fmt.Println()
	// 原来的长度和容量都为3，在使用append对切片进行扩容时，原来的切片容量不够，从新开辟出新的内存空间，将原来的切片容量翻倍，此时容量为6
	s1 = append(s1,1)
	fmt.Printf("s1的内存地址为%p",s1)
	fmt.Println()
	fmt.Println(s1)
	fmt.Println(cap(s1))
	fmt.Println(len(s1))
	//此时的内存地址与上面的一样，因为此时切片的长度并没有超出容量
	s1 = append(s1,2,3)
	fmt.Printf("s1的内存地址为%p",s1)
	fmt.Println()
	fmt.Println(s1)
	fmt.Println(cap(s1))
	fmt.Println(len(s1))
	//此时内存地址发生改变，容量发生改变，6*2=12
	s1 = append(s1,4)
	fmt.Printf("s1的内存地址为%p",s1)
	fmt.Println()
	fmt.Println(s1)
	fmt.Println(cap(s1))
	fmt.Println(len(s1))
}
