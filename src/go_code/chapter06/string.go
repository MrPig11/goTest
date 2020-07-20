package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	 s1 := "abcde"
	 //将字符串转为比特
	 sl := []byte(s1)
	 fmt.Println(sl)
	 //判断一个字符串是否存在另一个字符串中
	 is_ext := strings.Contains(s1,"ac")
	 fmt.Println(is_ext)
	 //判断一个字符串是否存在过另一个字符串中
	 is_ok := strings.ContainsAny(s1,"sdd")
	 fmt.Println(is_ok)
	 //判断一个字符串是否是以另一个字符串开头
	 s2 := strings.HasPrefix(s1,"a")
	 fmt.Println(s2)
	 //判断一个字符串是否以另一个字符串结尾
	 s2 = strings.HasSuffix(s1,"e")
	fmt.Println(s2)

	 fmt.Println("----------字符串与其他类型互相转换---------------")
	 i := 1
	 //整形转字符串
	 ss1 := strconv.Itoa(i);
	 fmt.Printf("%T%s\n",ss1,ss1)
	 //字符串转为整形
	 i,_ = strconv.Atoi(ss1)
	 fmt.Printf("%T%d",i,i)
	 //字符串转为bool
	 ss2 := "true"
	 b1,_ := strconv.ParseBool(ss2)
	 fmt.Printf("%T%t\n",b1,b1)
	 //布尔转为字符串
	 ss2 = strconv.FormatBool(b1)
	fmt.Printf("%T%s\n",ss2,ss2)
	 //字符串转为浮点
	 ss3 := "13.8"
	 f1,_ := strconv.ParseFloat(ss3,64)
	fmt.Printf("%T%f\n",f1,f1)
	 // 浮点转为字符串
	 ss3 = strconv.FormatFloat(f1,'e',1,32)
	fmt.Printf("%T%s\n",ss3,ss3)

}
