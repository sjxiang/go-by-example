package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// 用本地时间戳来初始化随机数种子
	rand.Seed(time.Now().UnixNano())
}

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)

	fmt.Println("请输入你猜的数字：")

	for {
		
		var guess int 
		_, err := fmt.Scanf("%d", &guess)

		if err != nil {
			fmt.Println("无效输入，请输入整数！")
			continue
		}

		fmt.Println("你的猜测是", guess)

		if guess > secretNumber {
			fmt.Println("不对，建议往小猜")
			continue
		} else if guess < secretNumber {
			fmt.Println("不对，建议往大猜")
			continue
		} else {
			fmt.Println("盖了帽了，我的老 baby！ 猜的真准")
			break
		}

	}
}

