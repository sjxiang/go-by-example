package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))                // truev 是否包含子字符串
	fmt.Println(strings.Count(a, "l"))                    // 2   统计出现频率
	fmt.Println(strings.HasPrefix(a, "he"))               // true  是否有相应的前缀
	fmt.Println(strings.HasSuffix(a, "llo"))              // true  
	fmt.Println(strings.Index(a, "ll"))                   // 2
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo
	fmt.Println(strings.Repeat(a, 2))                     // hellohello
	fmt.Println(strings.Replace(a, "l", "E", 2))          // heEEo  2 是替换数目，默认 -1
	fmt.Println(strings.Split("a-b-c", "-"))              // [a b c]
	fmt.Println(strings.ToLower(a))                       // hello
	fmt.Println(strings.ToUpper(a))                       // HELLO
	fmt.Println(len(a))                                   // 5
	b := "你好"
	fmt.Println(len(b)) // 6
}
