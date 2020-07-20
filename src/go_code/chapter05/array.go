package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 获取随机数,需要给随机数设置源，不然每次的随机数都一样
	a := rand.Intn(1000)
	fmt.Println(a)
	//源不能写死，可以使用时间戳,因为时间戳是int类型，所以要强转为int64类型
	rand.Seed(int64(time.Now().Second()))
	b := rand.Intn(1000)
	fmt.Println(b)
	// 时间戳的秒是int类型，纳秒是int64
	c := time.Now().UnixNano()
	fmt.Printf("%T",c)
fmt.Println()
	rand.Seed(time.Now().UnixNano())
	//获取一定范围的随机数 3-15
	d := rand.Intn(12)+3
	fmt.Println(d)
}
