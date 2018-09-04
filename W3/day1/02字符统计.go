package main

import (
	"io/ioutil"
	"fmt"
)

/*实现文件字符数统计*/
func CountCharsOfFile(path string)map[string]int {
	bytes, _ := ioutil.ReadFile(path)
	str := string(bytes)

	var numberCount,letterCount,spaceCount,othersCount int

	//遍历字符串中的每个字符
	for i,c := range str{
		fmt.Printf("No%d,%c\n",i,c)

		//逐个判断每个字符在字符集中的序号
		switch {
		case c >= '0' && c<='9':
			numberCount++
		case (c >= 'a' && c<='z') || (c >= 'A' && c<='Z'):
			letterCount++
		case c == ' ' || c=='\t':
			spaceCount++
		default:
			othersCount++
		}
	}
	return map[string]int{"数字":numberCount,"字母":letterCount,"空白":spaceCount,"其它":othersCount}
}
func main() {
	retMap := CountCharsOfFile("d:/temp/test.txt")
	fmt.Println(retMap)
}

func main0() {
/*	for _,c := range "fuck你妹off"{
		fmt.Printf("%T,%v,%U,%c\n",c,c,c,c)
	}*/
	fmt.Printf("%T,%v\n",'c','0')
	fmt.Printf("%T,%v\n",'c','1')
	fmt.Printf("%T,%v\n",'c','2')
	fmt.Printf("%T,%v\n",'c','3')
	fmt.Printf("%T,%v\n",'c','4')
	fmt.Printf("%T,%v\n",'c','5')
	fmt.Printf("%T,%v\n",'c','6')
	fmt.Printf("%T,%v\n",'c','7')
	fmt.Printf("%T,%v\n",'c','8')
	fmt.Printf("%T,%v\n",'c','9')
}

