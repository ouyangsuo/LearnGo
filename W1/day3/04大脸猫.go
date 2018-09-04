package main

import "fmt"

/*
标识符，用于命名常量、变量、表达式的字符
标识符可以是：大小写字母，数字，_
数字不能开头
*/

/*算算你的脸有多大*/

//定义浮点型常量PI
const PI float32 = 3.14

//定义整型变量radius，用于存储脸的半径
var radius int

//定义浮点型变量area，用于存储脸的面积
var area float32

//定义无参无返回值的函数
func demo41() {
	//给radius赋值为10（将10写入radius的内存地址）
	radius = 10

	//将radius强转为浮点型后，计算出脸的面积，赋值给area变量
	area = PI * float32(radius) * float32(radius)

	//重新给变量赋值
	//area = 108000

	fmt.Println("你的脸有", area, "辣么大")
	fmt.Printf("你的脸有%f辣么大\n", area)
	fmt.Printf("你的脸有%.2f辣么大\n", area)
	fmt.Printf("你的脸有%10.2f辣么大\n", area)

	radius = 5
	area = PI * float32(radius) * float32(radius)
	fmt.Printf("你的脸有%-10.2f辣么大\n", area)
}

//定义有参数函数getCircleArea
//参数：r用于接收半径
//驼峰风格的命名
func getCircleArea(r float32) {
	//var result float32
	//声明赋值合二为一，外加动态检测类型 变量名:=值
	result := PI * r * r
	fmt.Printf("result的类型为%T\n", result)
	fmt.Printf("半径为%.2f的脸有%.2f大\n", r, result)
}

//定义有参数、有返回值函数
//返回值为32位浮点
//小写加下划线风格
func get_circle_area(r float32) float32  {
	result := PI * r * r
	return result
}

//定义有参数、有返回值函数
//预定义了返回值的名字mianji
func get_circle_area_ii(radius float32) (mianji float32)  {

	//返回值变量mianji已经在函数定义中声明过了
	mianji = PI * radius * radius

	//已经定义了返回值的名字，无需在说明return谁
	return
}

func main() {
	//demo41()

	//getCircleArea(5)
	//getCircleArea(10)

	//接收函数的返回值，赋值给ret变量
	//ret := get_circle_area(5)
	ret := get_circle_area_ii(5)
	fmt.Printf("ret=%.2f",ret)

}
