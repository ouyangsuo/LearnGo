package main

import "fmt"

func main() {
	//if
	demo11()

	//switch
	demo12()

	//for
	demo13()

	//goto
	demo14()

}

func demo14() {
	fmt.Println("hello")
	fmt.Println("world")
	goto SHIT
	fmt.Println("领钱")
	fmt.Println("领美女")
SHIT:
	fmt.Println("GAME OVER")
}

func demo13() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//死循环
	for {
		fmt.Println("一起去浪")
	}
	//break跳出循环,continue跳过本次循环
	//打印1-100，逢5跳过，打到50就终止
	for i := 0; i < 100; i++ {
		if i > 50 {
			//跳出循环
			break
		} else if i%5 == 0 {
			//略过本次循环的剩余部分，开始下一次循环
			continue
		}
		fmt.Println(i)
	}
	//循环嵌套
	for i := 0; i < 10; i++ {
		//fmt.Printf("行%d",i)
		for j := 0; j < 10; j++ {
			fmt.Print("hello\t")
		}
	}
SHIT:
	for i := 0; i < 10; i++ {
		//fmt.Printf("行%d",i)
		for j := 0; j < 10; j++ {
			fmt.Print("hello\t")
			if i == 5 && j == 5 {
				break SHIT
			}
		}
	}
}

func demo12() {
	var level string
	//switch
	switch level {
	case "A":
		fmt.Println("super")
	case "B":
		fmt.Println("good")
	case "C":
		fmt.Println("normal")
	default:
		fmt.Println("bad")
	}
	//AB都OK
	switch level {
	case "A":
		fallthrough
	case "B":
		fmt.Println("good")
	case "C":
		fmt.Println("normal")
	default:
		fmt.Println("bad")
	}
}

func demo11() {
	if 3 > 2 {
		fmt.Println("yes")
	}
	if 3 > 2 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if 3 > 2 && 4 > 3 {
		fmt.Println("gt")
	} else if 3 == 2 || 4 == 3 {
		fmt.Println("eq")
	} else if !(5 == 3) {
		fmt.Println("other1")
	} else {
		fmt.Println("others")
	}
	var year int = 2018
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("is leap year")
	} else {
		fmt.Println("is not leap year")
	}
}
