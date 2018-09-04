package main

import "fmt"

/*标准输入输出*/

func getFaceArea(radius float32) (area float32) {
	area =  3.14 * radius * radius
	return
}

func getFaceArea2(name string, radius float32) {
	mianji :=  3.14 * radius * radius
	fmt.Printf("亲爱的%s：你有%.2f分薄面",name,mianji)
}

func main() {
	//demo51()
	demo52()
}

func demo51() {
	//让用户输入它的脸的半径
	fmt.Println("骚年请输入脸的半径：")
	var r float32
	//接收用户的输入，写入到r的地址中，&r取r的地址
	fmt.Scanf("%f", &r)
	mianji := getFaceArea(r)
	fmt.Printf("你的脸半径为%.2f,你的面子有%.2f", r, mianji)
}

func demo52() {
	var name string
	var radius float32

	fmt.Println("骚年请输入面部半径与姓名:")

	//接收用户的两次输入，以“+”为分隔符，分别写入到radius的地址、name的地址
	fmt.Scanf("%f+%s", &radius,&name)

	//将用户输入的name和radius作为参数，传给getFaceArea2函数进行处理
	getFaceArea2(name,radius)
}
