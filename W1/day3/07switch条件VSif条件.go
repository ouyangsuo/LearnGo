package main

import "fmt"

/*
if条件分支 VS switch条件分支
·if条件可以判断【连续的范围】，还支持条件与或非
·switch只适用于多个孤立的【点】，例如：1,2,3...或"a","b","c"...或true,false
*/

func main() {
	demo71()
}

func demo71() {
	var score int
	fmt.Println("骚年请输入期末考试成绩：")
	fmt.Scan(&score)

	//条件与：&&
	//条件或：||
	//条件表达式的值是一个连续的范围，所以switch不适用
	if score > 100 || score < 0 {
		fmt.Println("阅卷老师的水平已经远远超越了正常人类")
	} else if score >= 90 && score <= 100 {
		fmt.Println("您可以免费在兄弟连学习区块链，并到前台领取劳斯莱斯一辆")
	} else if score > 59 && score < 90 {
		fmt.Println("继续努力吧")
	} else {
		fmt.Println("区块链技术哪家强，山东找蓝翔")
	}

}
