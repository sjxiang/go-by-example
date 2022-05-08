package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "strings"
)

// 示例 {trans_type: "en2zh", source: "good"}

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source string `json:"source"`
	// UserID    string `json:"user_id"`
}

type DictResponse struct {
	Rc int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En string `json:"en"`
		} `json:"prons"`
		Explanations []string `json:"explanations"`
		Synonym []string `json:"synonym"`
		Antonym []string `json:"antonym"`
		WqxExample [][]string `json:"wqx_example"`
		Entry string `json:"entry"`
		Type string `json:"type"`
		Related []interface{} `json:"related"`
		Source string `json:"source"`
	} `json:"dictionary"`
}


func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: jiyiStar Word
jiyiStar fuck
		`)
		os.Exit(1)
	}
	
	word := os.Args[1]
	query(word)
}

func query(word string) {
	client := &http.Client{}
	var request = DictRequest{
		TransType: "en2zh",
		Source: word,
	}
	
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	// 所谓 "流" stream，网络传输 I/O 与 浏览器 v8 引擎渲染速率不匹配，需要将就下，读一点处理一点。
	// Reader s 初始化字符串 i 已读计数
    // [110 34 125]
	// &{[110 34 125] 0 -1}
	var data = bytes.NewReader(buf)

	// 创建请求
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求头
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭 stream
	defer resp.Body.Close()  

	// 读取响应
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)

	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}

	var dictresponse DictResponse
	err = json.Unmarshal(bodyText, &dictresponse)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%#v \n", dictresponse)

	fmt.Println(word, "UK:", dictresponse.Dictionary.Prons.En, "US:", dictresponse.Dictionary.Prons.EnUs)

	for _, item := range dictresponse.Dictionary.Explanations {
		fmt.Println(item)
	}

}



/*

1. 抓包 
	网页 - 右键菜单 检查 -  network


2. 代码生成 - HTTP 请求
	点击请求 - 右键菜单 copy - copy as cURL(bash)

	https://curlconverter.com/#go 


3. 生成代码解读

	创建请求
	设置请求头
	发送请求
	读取响应


4. 生成 request body


5. 解析 response body

	JSON 转 struct
	
	https://oktools.net/json2go


6. 完善代码




作业：
1. 增加另一种翻译引擎的支持
2. 在 "1" 基础之上，并行请求两个翻译引擎来提高响应速度


*/
