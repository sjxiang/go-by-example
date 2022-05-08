package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	// "bytes"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "usage: jiyiStar hello")
		os.Exit(1)
	}

	fmt.Println("双核高并发！")

	var wg sync.WaitGroup
	wg.Add(2)

	word := os.Args[1]

	go func() {
		defer wg.Done()
		queryHuoShan(word)
	}()

	go func() {
		defer wg.Done()
		queryCaiyun(word)
	}()
	
	wg.Wait()

}


func queryCaiyun(word string) {
	fmt.Println("彩云翻译为您提供服务")

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
	fmt.Println()
}


func queryHuoShan(word string) {
	fmt.Println("火山翻译为您提供服务")

	client := &http.Client{}

	var huoReq = huoshanFanyiRequest{
		Text: word + "\n",
		Language: "en",
	}

	buf, err := json.Marshal(huoReq)
	var data = bytes.NewReader(buf)
	
	req, err := http.NewRequest("POST", "https://translate.volcengine.com/web/dict/match/v1/?msToken=&X-Bogus=DFSzKwGLQDVBKogvSW/mL-t/pLvl&_signature=_02B4Z6wo000015vBRYQAAIDCEIreZG4oMxOb0UEAAISPxhIbxUIuVNG4hYvQAvYEBiwvJ8WiLiX1YCVwfpArYAvhcBJbWVgF8sRwtEkyO5gmHFvd6NG.fNSUcA0Hgd0o-.QESIm4xFwoqJBK79", data)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("authority", "translate.volcengine.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "x-jupiter-uuid=16519932556812288; i18next=zh-CN; ttcid=17287189e4014ac6ba87bdfa6c04052712; tt_scid=y4Qq.GV2q7ag0iwIkfWhl2U2irTu1k8l55aCYRrLmINS-c70.j.6xZ5o3czIdGyy5db9; s_v_web_id=verify_06240da56eac839d85bf8dd0810b2945; _tea_utm_cache_2018=undefined")
	req.Header.Set("origin", "https://translate.volcengine.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://translate.volcengine.com/translate?category=&home_language=zh&source_language=detect&target_language=zh&text=good%0A")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Microsoft Edge";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", `"Android"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Mobile Safari/537.36 Edg/101.0.1210.39")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	
	var huoResp huoshanFanyiResponse
	err = json.Unmarshal(bodyText, &huoResp)
	
		// words[0].pos_list[0].explanations[0]
		// words[0].pos_list[1].explanations[0]

	// fmt.Println(huoResp.Words[0].PosList[0].Explanations[0].Text, huoResp.Words[0].PosList[1].Explanations[0].Text)// words[0].pos_list[0].explanations  words[0].pos_list[0].explanations[0]
	fmt.Println(word, "tips：音标省略，调字段太烦 😔 ")

	for _, v := range huoResp.Words[0].PosList {
		fmt.Println(v.Explanations[0].Text)
	}

	fmt.Println()
}