package main

import (
	"fmt"
	"sort"
)

func main() {
	//创建map数据类型 默认为nil，不能直接使用
	var map1  map[int] string
	fmt.Println(map1 == nil)
	map2 := make(map[string] int)
	fmt.Println(map2)
	var map3 = map[int] string{1:"哈哈",2:"hello go",3:"张三",4:"李四",5:"王五"}
	fmt.Println(map3)

	// 获取map数据
	value,ok := map3[1]
	if ok {
		//为true，说明获取的是map的值
		fmt.Println("我获取到map的值了",value)
	} else {
		fmt.Println("我获取到的是map的默认值")
	}
	//删除map数据
	//delete(map3,1)
	//fmt.Println(map3)
	//fmt.Println(len(map3))
	//map4 := map3
	//map4[2] = "修改一下"
	//此时map4与map3相等，所以map是址传递
	//fmt.Println(map4)
	//fmt.Println(map3)
	//遍历map ，map是无序的，打印出来每次的顺序都不一样，使用切片
	//for i,v := range map3 {
	//	fmt.Println(v,i)
	//}
	//创建一个切片
	s1 := make([]int,0,len(map3))
	for i,_ := range map3{
		s1 = append(s1,i)
	}

	//进行排序，按照升序进行排列
	sort.Ints(s1)
	fmt.Println(s1)
	for a := 1;a <= len(map3);a++{
		fmt.Println(map3[a])
	}
}
