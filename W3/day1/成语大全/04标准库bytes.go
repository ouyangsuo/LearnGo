package main

import (
	"fmt"
	"bytes"
	"strings"
)

func main() {
	/*
	func Title(s []byte) []byte
	func ToLower(s []byte) []byte
	func ToUpper(s []byte) []byte
	*/

	fmt.Println(string(bytes.ToLower([]byte("ShIt"))))
	fmt.Println(string(bytes.ToUpper([]byte("ShIt"))))
	fmt.Println(string(bytes.Title([]byte("sHiT"))))

	//strs := []string{"hello", "world"}
	//bytess := [] []byte {[]byte("hello"), []byte("world")}
	//fmt.Println(bytess)
	//fmt.Println(bytess[1][2])
	fmt.Println('d')//100

	fmt.Println(strings.Split("hello,world,shitEater",","))//[hello world shitEater]
	fmt.Println(bytes.Split([]byte("hello,world,shitEater"),[]byte(",")))//[[1 2 3 4 5] [6 5 4 3 2 1] [6 6 6 6 6 6 6 6]]

	//创建一江水的读取器
	reader := bytes.NewReader([]byte("你妹hello,world,shitEater"))
	//创建2字节的小水桶作为缓冲区
	buffer := make([]byte, 8)
	//从江中打一桶水
	n, _ := reader.Read(buffer)
	//打印出来看看
	fmt.Println(n,string(buffer))//8 你妹he 这桶水的长度是8字节 内容是“你妹he”（每个ascii字符占1字节，每中文3字节）

	//从江中打第二桶水
	n,_ = reader.Read(buffer)
	fmt.Println(n,string(buffer))//8 llo,worl

}
