package main

import "fmt"

func main() {
	//创建map，存储个人信息
	map1 := map[string]string{"name":"张三","age":"12","sex":"男"}
	map2 := map[string]string{"name":"李四","age":"17","sex":"男"}
	map3 := map[string]string{"name":"王五","age":"19","sex":"女"}
	//将map存入切片中，声明存入切片中的map的数据类型
	s1 := make([]map[string]string,0,3)
	s1 = append(s1,map1)
	s1 = append(s1,map2)
	s1 = append(s1,map3)
	fmt.Println(s1)
	for k, v := range s1 {
		fmt.Printf("键是%d",k)
		fmt.Printf("值是%s",v["name"])
		fmt.Printf("值是%s",v["age"])
		fmt.Printf("值是%s",v["sex"])
		fmt.Println()
	}
	//创建数组，二维数组
	arr1 := [4][4]int{{1,2,3,4},{1,2,3,4},{1,2,3,4},{1,2,3,4}}
	fmt.Println(arr1)
}
