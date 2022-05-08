package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	
	// 只读的 stream
	reader := bufio.NewReader(os.Stdin)  // 程序执行 一般 有三个文件 

	for {

		// 读取 1 行
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入时，发生了错误，请重新输入", err)
			continue 
		}
	
		// 去除字符串后缀
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)

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



/*

1. 生成随机数
2. 读取用户输入
3. 判断逻辑
4. 实现游戏循环


作业：
	fmt.Scanf 简化代码


	var guess int
	_, err := fmt.Scanf("%d", &guess)

	
*/