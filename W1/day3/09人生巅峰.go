package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("亲请输入下注金额和下注球队:")

	var money int
	var team string
	fmt.Scanf("%d+%s", &money, &team)
	fmt.Printf("您下注了%s%d万元，人生巅峰即将开始...\n",team,money)
	time.Sleep(1 * time.Second)

	//time.Now().Unix()获取当前时间距离1970年零时逝去的秒数
	//rand.NewSource(time.Now().Unix())每秒更新一个随机数的种子,一旦种子变化随机数也随之变化
	myrand := rand.New(rand.NewSource(time.Now().Unix()))

	//获得[0-100)的随机数
	luckyNumber := myrand.Intn(100)
	fmt.Println("luckyNumber=", luckyNumber)

	if luckyNumber > 10 {
		fmt.Println("靠海别野欢迎你")
	} else {
		fmt.Println("天台欢迎你")
	}

}
