package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:9000") // 修改为你的服务器地址和端口
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return
	}
	defer conn.Close()

	// 使用 bufio.Reader 读取输入
	reader := bufio.NewReader(os.Stdin)

	// 发送用户 ID
	fmt.Print("请输入用户 ID: ")
	userID, _ := reader.ReadString('\n')
	userID = strings.TrimSpace(userID)
	fmt.Fprintf(conn, "user:%s\n", userID)

	// 创建一个 channel 来接收消息
	messageChannel := make(chan string)

	// 启动一个 goroutine 用于接收消息
	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("连接已关闭:", err)
				return
			}
			messageChannel <- message
		}
	}()

	// 启动一个 goroutine 用于发送消息
	go func() {
		for {
			fmt.Print("✅ 请输入聊天消息 (格式: talk:toUserID:message): ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)
			if message != "" {
				fmt.Fprintf(conn, "%s\n", message) // 发送消息时添加换行符
			}
		}
	}()

	// 主循环，接收消息
	for msg := range messageChannel {
		fmt.Println() // 打印空行
		fmt.Print("🚀 收到消息: ", msg)
		fmt.Print("✅ 请输入聊天消息 (格式: talk:toUserID:message): ") // 再次打印输入提示
	}
}
