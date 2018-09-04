package game

import (
	"fmt"
	"strconv"
)

//定义人员结构体
type person struct {
	name  string
	money int
}

func sortPersonArray(ps []person) {
	for i:=0;i<len(ps)-1;i++{
		for j:=i+1;j<len(ps);j++{
			if ps[j].money > ps[i].money{
				ps[i],ps[j] = ps[j],ps[i]
			}
		}
	}
}

func StartForbs() {
	//人员切片容器
	var ps []person
	//循环接收用户的录入
	for {
		var name, money string

		//接收用户名
		fmt.Print("请输入姓名：")
		fmt.Scanf("%s", &name)
		if name == "over" {
			break
		}

		//接收财富
		fmt.Print("请输入财富：")
		fmt.Scanf("%s", &money)
		if money == "over" {
			break
		}

		//将字符串money转换为int
		ret, _ := strconv.Atoi(money)

		//造一个person对象
		p := person{name, ret}

		//将人员对象p丢入人员容器ps
		ps = append(ps, p)
	}
	//使用选择排序法按财富排序人员
	sortPersonArray(ps)
	//输入出人信息
	fmt.Println(ps)
}
