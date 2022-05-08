package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	// 连接的生命周期 即函数的生命周期
	defer conn.Close()

	// 基于连接创建一个带缓冲的 stream
	reader := bufio.NewReader(conn)
	
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		
		fmt.Printf("%T", b)  // uint8

		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}




/*

socks5 代理

明文传输

$ curl --socks5 127.0.0.1:1080 -v http://www.qq.com



正常访问网站
	tcp 3 次握手
	发送 HTTP 请求
	返回 HTTP 响应



client      socks5 代理服务器         目标服务器
	  1. 协商
	  1.1 通过协商

	  2 发送请求              2.1 建立TCP 连接
      2.3 返回状态            2.2 返回响应

	  3. 发送数据             3.1 relay 数据
	  3.3 响应结果            3.2 响应结果

	认证
	请求


tcp echo server
$ nc 127.0.0.1 1080  // 连接远程



SwitchyOmega


https://segmentfault.com/a/1190000015591319#:~:text=io.Copy%20%28%29%20%E5%8F%AF%E4%BB%A5%E8%BD%BB%E6%9D%BE%E5%9C%B0%E5%B0%86%E6%95%B0%E6%8D%AE%E4%BB%8E%E4%B8%80%E4%B8%AA%20Reader%20%E6%8B%B7%E8%B4%9D%E5%88%B0%E5%8F%A6%E4%B8%80%E4%B8%AA%20Writer%E3%80%82,%E5%AE%83%E6%8A%BD%E8%B1%A1%E5%87%BA%20for%20%E5%BE%AA%E7%8E%AF%E6%A8%A1%E5%BC%8F%EF%BC%88%E6%88%91%E4%BB%AC%E4%B8%8A%E9%9D%A2%E5%B7%B2%E7%BB%8F%E5%AE%9E%E7%8E%B0%E4%BA%86%EF%BC%89%E5%B9%B6%E6%AD%A3%E7%A1%AE%E5%A4%84%E7%90%86%20io.EOF%20%E5%92%8C%20%E5%AD%97%E8%8A%82%E8%AE%A1%E6%95%B0%E3%80%82

https://segmentfault.com/blog/sxsgo

*/