package server

import (
	"fmt"
	"net"
	"sync"
)

// 存储用户 ID 和 TCP 连接的映射
var userConnMap = make(map[string]net.Conn)
var mapMutex sync.Mutex

// 启动TCP服务器
func StartTCPServer() {
	tcpListen, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("启动TCP服务器失败:", err)
		return
	}
	defer tcpListen.Close()

	fmt.Println("TCP服务器已启动，等待连接...")

	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			fmt.Println("接收连接失败:", err)
			continue
		}

		// 使用协程处理每个连接
		go handleConnection(conn)
	}
}
