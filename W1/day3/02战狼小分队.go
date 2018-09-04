package main


/*import (
	//系统的格式化包
	"fmt"
	//引入系统时间包
	"time"
)*/

import "fmt"
import "time"

/*
单协程内部，自上而下【串行】执行
多协程，【并行】执行（同时执行）
*/

//将在独立的协程中并发执行
func doSomething(i int){
	//如果子协程睡眠的过程中，主协程结束了，那么子协程会提前结束
	//这里的代码是跑在子协程中
	time.Sleep(1*time.Second)
	fmt.Println("我是战狼小分队",i)
}

func main() {

	//main中的代码都是跑在主协程中
	fmt.Println("总部要行动了!")

	//其它语言并发：进程-线程（并发的资源开销大）
	//并发100条【协程】（微线程）
	for i:=1;i<100;i++{

		//鉴赏一下Go开并发是夺妹地容易...
		//开辟独立的子【协程】中并发执行doSomething(i)
		go doSomething(i)
	}

	fmt.Println("总司令晕过去了...")
	//为什么要让主协程睡觉？因为主协程如果结束，子协程都会结束
	time.Sleep(2*time.Second)
	fmt.Println("GAME OVER!")

}
