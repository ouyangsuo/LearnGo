package main

import "fmt"

func main() {
	//fmt.Println(add4(3, 4, 5, 6))
	//fmt.Println(add5("add", true, 5, 6, 7))
	add6("add", true, 5, 6.5, "fuckoff",false)
}

func add1(a int, b int) int {
	return a + b
}

//a,b都是整型，一次性声明类型
func add2(a, b int) int {
	return a + b
}

//定义返回值名字
func add3(a, b int) (sum int) {
	sum = a + b
	return
}

//变长参数
func add4(a ... int) (sum int) {
	fmt.Printf("type=%T,value=%v\n", a, a)
	//简易遍历数据容器,index为序号，value为值
	for index, value := range a {
		fmt.Println(index, value)
		sum += value
	}
	return
}

//变长参数
func add5(arg1 string, arg2 bool, a ... int) (sum int) {
	fmt.Println(arg1, arg2, a)
	fmt.Printf("type=%T,value=%v\n", a, a)
	//简易遍历数据容器,index为序号，value为值
	for index, value := range a {
		fmt.Println(index, value)
		sum += value
	}
	return
}

//变长参数且类型也不定
func add6(arg1 string, arg2 bool, a ... interface{}) {
	fmt.Println(arg1, arg2, a)
	fmt.Printf("type=%T,value=%v\n", a, a)
	//简易遍历数据容器,index为序号，value为值
	for index, value := range a {
		fmt.Println(index, value)
	}
}

//定义返回值名字
func calc1(a, b int) (int, float32) {
	sum := a + b
	//int和int求商还是int，需要强转一下
	shang := float32(a) / float32(b)
	return sum, shang
}

func calc2(a, b int) (sum int, shang float32) {
	sum = a + b
	//int和int求商还是int，需要强转一下
	shang = float32(a) / float32(b)
	return
}

func calculate(a, b int, operator string) float32 {
	if operator == "+" {
		return float32(a) + float32(b)
	}
	if operator == "%" {
		return float32(a % b)
	}

	return -1.0
}
