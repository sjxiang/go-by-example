package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// go run example/20-env/main.go a b c d
	fmt.Println(os.Args[1:])           // [ a b c d]
	fmt.Println(os.Getenv("PATH"))     // /usr/local/go/bin...
	fmt.Println(os.Setenv("AA", "BB")) // 设置环境变量

	buf, err := exec.Command("curl", "www.baidu.com").CombinedOutput()  // 执行 shell 获得输入输出
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf)) 
}
