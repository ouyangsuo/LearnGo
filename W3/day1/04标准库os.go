package main

import (
"os"
"fmt"
"time"
)

func main() {
	//D:\BJBlockChain1801\demos\
	dir, _ := os.Getwd()
	fmt.Println(dir)

	//D:\iWorkspace\GoPros\Go18DaysCode\Day13project\;C:\Users\sirouyang\go;D:\BJGo1801Pre\preWorks\predemos\W99\03标准库\38单元测试
	paths := os.Getenv("GOPATH")
	fmt.Println(paths)

	os.Chtimes("d:/temp/小黑子.avi",time.Now(),time.Now())

	environ := os.Environ()
	fmt.Println(environ)

	fmt.Println(os.Hostname())

	fmt.Println(os.IsPathSeparator('/'))//Linux认
	fmt.Println(os.IsPathSeparator('\\'))//Linux不认

	fileInfo1, _ := os.Stat("d:/temp/小黑子2.avi")
	fileInfo2, _ := os.Stat("d:/temp/小黑子"+"2"+".avi")
	fmt.Println(os.SameFile(fileInfo1,fileInfo2))

	fmt.Println(os.TempDir())

}
