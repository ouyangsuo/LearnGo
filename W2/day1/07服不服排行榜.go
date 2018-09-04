package main

import "game"

/*
·循环让用户输入人员信息，首先输入用户名，回车后提示输入财富金额，如此循环往复；
·当用户输入over时，结束输入，并输入财富排行榜，包括姓名和财富；
·将这个游戏封装在game包下的Forbs()方法中，在入口函数中进行调用；
*/

func main() {
	game.StartForbs()
}
