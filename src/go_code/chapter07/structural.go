package main

import (
	"fmt"
)

func main() {
	p1 := Person{}
	p1.name = "张三"
	p1.age = 18
	p1.sex = "男"
	p1.address = "我是地址"
	fmt.Printf("p1的姓名是:%s,年龄是:%d,性别是:%s，地市是:%s\n",p1.name,p1.age,p1.sex,p1.address)

	p2 := Person{name:"李四",age: 13,sex:"女",address: "我是李四的地址"}
	fmt.Println(p2)

	//结构体指针
	p3 := new(Person)
	p3.name = "hahah"
	//都是指针
	p4 := p3
	p4.name = "啊啊啊"
	fmt.Println(p3)
	fmt.Println(p4)
	var p5 *Person
	//取出p1的内存地址  赋值给p5
	p5 = &p1
	//fmt.Println(p5)
	//fmt.Println(pp1)
	p5.name = "ppp"
	fmt.Println(p5)


	p6 := Book{"雪山飞狐",98.6}
	//结构体的嵌套
	p7 := Student{name: "王五",age:19,book:Book{"hahah",87.7}}
	fmt.Println(p6)
	fmt.Println(p7)
	//传递指针
	p8 := Student2{name: "王五",age:19,book:&p6}
	fmt.Println(p8)

	p8.book.bookName = "笑傲江湖"
	p8.book.price = 16.9
	fmt.Println(*p8.book)


	p9 := Student3{}
	p9.name = "awwww"
	fmt.Println(p9)
	//eay是student2中的方法，但是子类也可以使用
	p9.eay()
}

//结构体
type Person struct {
	name string
	age int
	sex string
	address string
}
type Book struct {
	bookName string
	price float64
}

type Student struct {
	name string
	age int
	book Book
}

type Student2 struct {
	name string
	age int
	//传递指针
	book *Book
}

type Student3 struct {
	//结构体的继承
	Student2
}
//方法，指定谁去使用
func (s Student2) eay(){
	fmt.Println("我要学习")
}
//重写方法，方法的覆盖
func (s Student3)eay()  {
	fmt.Println("我从写的student2的方法")
}