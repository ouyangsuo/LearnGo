package main

import "game"

/*
·猜数字游戏：系统随机生成1000以内的随机数，循环让用户输入答案，系统反馈输大了输小了还是猜对了，直到猜对为止；
·如果用户输入“退散”提前终止游戏并公布答案；
·将这个游戏封装在game包下，通过GuessNumber()方法进行调用；
*/

func main() {
	game.GuessNumber()
}


