package yourmath

import (
	"fmt"
	"math"
)

//算法说明：
//以27为例，循环判断[2-26]之间有没有27的因子，循环(27-2)次
func IsPrimeNumber(n int) bool {
	if n < 2{
		fmt.Println("请输入一个2以上的整数")
		return false
	}

	if n == 2{
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println("发现因子",i)
			return false
		}
	}
	return true
}

//优化算法说明：
//如果有因子，要么是①平方根*平方根②一个小于平方根的因子*一个大于平方根的因子
// 由于是成对出现，所以只需要循环到平方根处几个
//以27为例，循环判断[2-5]之间有没有27的因子，循环(27的平方根-1)次
func IsPrimeNumber2(n int) bool {
	if n < 2{
		fmt.Println("请输入一个2以上的整数")
		return false
	}

	if n == 2{
		return true
	}

	squereRoot := math.Sqrt(float64(n))
	for i := 2; i < int(squereRoot)+1; i++ {
		if n%i == 0 {
			fmt.Println("发现因子",i)
			return false
		}
	}
	return true
}
