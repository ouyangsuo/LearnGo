package main

import (
	"io"
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
)

/*
·实现两种形式的图片拷贝
----------
·实现对任意文档的字符统计（字母数、空白数、数字数、其它字符）
*/

func main11() {
	srcFile, err := os.OpenFile("d:/temp/fuckoff.jpg", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("源文件打开失败,err=", err)
	}
	dstFile, err := os.OpenFile("d:/temp/fuckoff2.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	defer srcFile.Close()
	defer dstFile.Close()

	if err != nil {
		fmt.Println("目标文件打开失败,err=", err)
	}
	written, err := io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("拷贝失败,err=", err)
		return
	}
	fmt.Println("拷贝成功！字节数=", written)
}

func main12(){
	bytes, _ := ioutil.ReadFile("d:/temp/fuckoff.jpg")
	err := ioutil.WriteFile("d:/temp/fuckoff2.jpg", bytes, 0666)
	if err != nil {
		fmt.Println("拷贝失败,err=", err)
		return
	}
	fmt.Println("拷贝成功！")
}

func main13() {
	//打开源文件和目标文件，并即刻挂起关闭程序
	srcFile, _ := os.OpenFile("d:/temp/fuckoff.jpg", os.O_RDONLY, 0666)
	dstFile, _ := os.OpenFile("d:/temp/fuckoff2.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	defer srcFile.Close()
	defer dstFile.Close()

	//创建读写器
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)

	//循环IO
	for {
		//创建4字节的缓冲区
		buffer := make([]byte, 4)

		//读入4个字节的数据到缓冲区
		n, err := reader.Read(buffer)

		//如果已经到了文件末尾，且最后一桶数据（buffer）为空
		//读取完毕，结束循环
		if n == 0 && err == io.EOF{
			break
		}

		//将缓冲区中的数据“倒入”目标文件
		writer.Write(buffer)

		//清空缓冲区
		writer.Flush()
	}
	fmt.Println("拷贝成功！")

}
