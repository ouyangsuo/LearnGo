package main

import "fmt"

//constant常量 variation变量 function函数 print打印 line行
const pi float32 = 3.14

//动态判别类型
const LightSpeed = 300000
var isClever = false

//一次性声明多个
const(
	Chidao = 40000
	ChinaArea int = 960
)
var(
	fug = "hello"
	shid int = 123
)

//声明布尔型变量
//var isStupid bool = true//真
//var isStupid bool = false//假
var isStupid bool

/*声明数值型变量*/
//var age int = 60
var age int
//var rmb float32 = 0.5
var rmb float32
//复数
//var vector complex64 = 3+4i
var vector complex64

//声明字符串变量
//var name string = "张全蛋"
var name string

/*声明复合型变量*/
//10长度的整型数组
//var ages [10]int = [10]int{1,2,3,4,5,6,7,8,9,0}
var ages [10]int

//可变长的整型切片
//var heights []int = []int{1,2,3}
var heights []int

//映射（键1：值1，键2：值2...）
//var kpi map[string]int = map[string]int{"代码量":10000,"注释量":3000}
var kpi map[string]int

var weight int = 300
//整型指针，weightPointer存放的不是300，而是放置300的内存地址
//var weightPointer *int = &weight
var weightPointer *int

//定义公有的（首字母大写，否则包内私有）接口Animal
//Go语言是强类型语言，但是空接口interface{}可以代指任何类型
var Animal interface{}



//结构体
//var person struct {
//	name string
//	age int
//	rmb float32
//}
var person struct{}

//声明函数变量
//var eat func() = func() {
//	fmt.Println("我是个吃货")
//}
var eat func(a int,b float64)bool

func main() {

	fmt.Println("常量LightSpeed的实际类型是%T\n",LightSpeed)
	fmt.Println("变量isClever的实际类型是%T\n",isClever)

	//变量的声明赋值合二为一，这种声明方式只能在函数内部
	dann := 456.78
	fmt.Printf("type=%T,value=%v\n",dann,dann)

	/*
	fmt.Printf格式化打印
	%T 类型占位符
	%v 值占位符
	*/
	fmt.Printf("【是否愚蠢】的类型是%T，值是%v\n",isStupid,isStupid)
	fmt.Printf("【年龄】的类型是%T，值是%v\n",age,age)
	fmt.Printf("【资产】的类型是%T，值是%v\n",rmb,rmb)
	fmt.Printf("【向量】的类型是%T，值是%v\n",vector,vector)
	fmt.Printf("【姓名】的类型是%T，值是%v\n",name,name)
	fmt.Printf("【年龄们】的类型是%T，值是%v\n",ages,ages)
	fmt.Printf("【身高们】的类型是%T，值是%v\n",heights,heights)
	fmt.Printf("【KPI】的类型是%T，值是%v\n", kpi, kpi)
	fmt.Printf("【存体重的地址】的类型是%T，值是%v\n",weightPointer,weightPointer)
	fmt.Printf("【人】的类型是%T，值是%v\n", person, person)
	fmt.Printf("【饕餮】的类型是%T，值是%v\n",eat,eat)

}