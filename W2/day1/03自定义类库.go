package main

import (
	//引入系统包
	"fmt"
	"time"

	//引入自定义包
	//自定义的包只有注册在GOPATH下才能使用
	//注册方式：将【yourmath包所在的src目录的上层目录】写入系统的GOPATH环境变量中
	"yourutils/yourmath"
)

func main() {
	fmt.Println("使用自定的类库")
	fmt.Println("当前时间",time.Now())
	ret := yourmath.Subtract(3,4)

	fmt.Println(ret)
}
