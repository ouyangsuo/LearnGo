package main

import "fmt"

type Person struct {
	name string
	age  int
}

//结构体的默认值是空白结构体（属性都是默认值）
//指针的默认值是nil
func main1() {
	var p Person
	var pp *Person

	fmt.Println(p)
	fmt.Println(pp)
}

//
func main() {
	pMap := make(map[string]Person)
	//键不存在时，值是空白对象（属性都是默认值）
	value, ok := pMap["张全蛋"]
	fmt.Println(value, ok)

	ppMap := make(map[string]*Person)
	//键不存在时，值是nil
	vp, ok := ppMap["张全蛋"]
	fmt.Println(vp, ok)
}
