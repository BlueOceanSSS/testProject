package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:8898")
	defer dial.Close()
	if err != nil {
		fmt.Println("连接失败")
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("连接失败")
			os.Exit(1)
		}
		dial.Write(line)
		fmt.Println("发送消息")

		n, err2 := dial.Read(buffer)
		if err2 != nil {
			fmt.Println("连接失败")
			os.Exit(1)
		}
		msg := string(buffer[:n])
		fmt.Println("客户端收到消息：", msg)

	}
}
