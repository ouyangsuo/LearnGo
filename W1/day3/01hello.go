//劳资是单行注释，快捷键是ctrl+/

/*
劳资是段落注释
快捷键是ctrl+shift+/
*/

/*
//生成可执行程序01hello.exe
go build ./01hello.go

//指定输出位置为./fuck.exe
go build -o ./fuck.exe ./01hello.go

//执行编译好的程序
./01hello.exe

//编译执行合二为一
go run ./01hello.go
*/

//包名为main，入口程序的包名必须是main
package main

//引入系统包(系统提供的支持库)
import "fmt"

//入口函数，名字必须是main
func main(){

	//调用系统包fmt下的打印函数，打印内容
	fmt.Println("Hello,World!Let's Golang!")
	fmt.Println("师姐你好，让我们去浪吧！")

}
