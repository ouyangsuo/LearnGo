package main

import (
	"fmt"
	"time"
)

func main() {
	////demo81()
	//demo82()
	//demo83()

	//当条件是孤立的点时，if和switch都可以
	demo84()
	demo85()
}

//爱阿根-爱战车-最爱是天台
func demo84() {
	for i := 1; i < 11; i++ {
		if i%5 == 0 {
			fmt.Println("最爱是天台")
		} else if i%2 == 1 {
			fmt.Println("爱阿根廷")
		} else {
			fmt.Println("更爱战车")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

//爱要强-爱铁血-最爱是天台
func demo85() {

	for i := 1; i < 11; i++ {
		switch i % 5 {

		//1,3都走爱要强
		case 1:
			//继续向下执行
			fallthrough
		case 3:
			fmt.Println("爱要强")

			//2,4都走爱战车
		case 2:
			fallthrough
		case 4:
			fmt.Println("爱铁血")

			//0
		default:
			fmt.Println("最爱是天台")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

//有限次循环
func demo83() {
	//起始条件：i=1，循环条件：i<11，增长条件：i++
	for i := 1; i < 11; i++ {
		fmt.Println("爱天台，爱战车", i)
		time.Sleep(500 * time.Millisecond)
	}
}

//自增与自减
func demo82() {
	var i int = 0

	//i+1的结果，重新赋值给i，等价于i++，自增
	i = i + 1
	fmt.Println(i)
	//1
	i++
	fmt.Println(i)
	//2
	i--
	fmt.Println(i)
	//1
}

func demo81() {
	//无限死循环
	for {
		fmt.Println("爱天台，爱日耳曼战车")
		time.Sleep(1 * time.Second)
	}
}
