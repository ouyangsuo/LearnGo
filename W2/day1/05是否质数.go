package main

import (
	"fmt"
	"yourutils/yourmath"
)

/*
//7 (2-6)
·在myutil/mymathutil的包下继续封装函，判断输入的整数参数是否是质数（因子只有1和本身）
·在入口函数中对封装的函数进行测试；
*/

func main() {

	var guess int
	for {
		fmt.Scanf("%d\n",&guess)
		fmt.Printf("%t\n", yourmath.IsPrimeNumber2(guess))
	}
}
