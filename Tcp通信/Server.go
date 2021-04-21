package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8898")
	defer listen.Close()
	handleErr(err)
	for {
		fmt.Println("开始循环接收连接...")
		accept, err := listen.Accept()
		handleErr(err)
		go func(accept net.Conn) {
			buf := make([]byte, 1024)
			defer accept.Close()
			for {
				numOfBytes, err2 := accept.Read(buf)
				handleErr(err2)
				if numOfBytes != 0 {
					msg := string(buf[:numOfBytes])
					fmt.Println("Server：收到消息", msg)
					accept.Write([]byte("已阅" + msg))
					if msg == "断开" {
						break
					}
				}
			}

		}(accept)
	}

}

func handleErr(err error) {
	if err != nil {
		fmt.Println("程序出现错误")
		os.Exit(1)
	}
}
