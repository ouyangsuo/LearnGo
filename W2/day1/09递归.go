package main

import "fmt"

/*
递归：自己调用自己
·要点：一定要有递归终止条件
*/

func main() {
	//demo91()
	//fmt.Println(getSum(10))
	for i:=0;i<100 ;i++  {
		fmt.Println(i,getFibonacci(i))
	}

}

//1+2+...+10的常规实现
func demo91() {
	var sum = 0
	for i := 1; i < 11; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//1+2+...+10的递归(recursive)实现
func getSum(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + getSum(n-1)
	}
}

//斐波那契数列：1,1,2,3,5,8,13,21...
func getFibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return getFibonacci(n-1) + getFibonacci(n-2)
	}
}
