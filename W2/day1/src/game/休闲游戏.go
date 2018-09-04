package game

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

func GuessNumber() {
	randPoint := rand.New(rand.NewSource(time.Now().UnixNano()))
	answer := randPoint.Intn(1000)
	var guess string
	for {
		fmt.Print("大官人请输入你的猜想：")
		fmt.Scanf("%s", &guess)

		if guess == "退散" {
			fmt.Printf("answer=%d\n", answer)
			break
		}

		//大小判断
		ret, _ := strconv.Atoi(guess)
		if ret > answer {
			fmt.Println("猜大了，垃圾")
		} else if ret < answer {
			fmt.Println("猜小了，乐色")
		} else {
			fmt.Println("BINGO")
			break
		}
	}
	fmt.Println("GAME OVER")
}
