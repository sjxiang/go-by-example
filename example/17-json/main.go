package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {    
	Name  string               
	Age   int `json:"age"`  // 保证 "对应字段" 大写，即可
	Hobby []string
}

func main() {
	a := userInfo{Name: "相", Age: 18, Hobby: []string{"Golang", "Python"}}
	buf, err := json.Marshal(a)  // 序列化输出 byte 切片
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)          // [123 34 78 97...]
	fmt.Printf("%T\n", buf)   // []uint8 也就是 byte 切片
	fmt.Println(string(buf))  // {"Name":"相","age":18,"Hobby":["Golang","Python"]}

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)  // 反序列化到 empty 结构体
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", b) // main.userInfo{Name:"wang", Age:18, Hobby:[]string{"Golang", "TypeScript"}}
}
