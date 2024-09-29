package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// 处理用户 TCP 连接
func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	var userID string

	for {
		// 读取消息
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("用户 %s 断开连接\n", userID)
			// 删除失效连接
			if userID != "" {
				removeConnection(userID)
			}
			return
		}
		message = strings.TrimSpace(message)

		// 处理用户 ID
		if strings.HasPrefix(message, "user:") {
			userID = strings.TrimSpace(strings.TrimPrefix(message, "user:"))
			if userID != "" {
				mapMutex.Lock()
				userConnMap[userID] = conn
				mapMutex.Unlock()
				fmt.Printf("用户 %s 已连接\n", userID)
				// 继续等待其他消息
				continue
			}
		}

		if strings.HasPrefix(message, "talk:") {
			// 假设完整消息格式为 "talk:fromUserID:toUserID:message"
			talkParts := strings.SplitN(strings.TrimPrefix(message, "talk:"), ":", 3)
			if len(talkParts) == 3 {
				fromUserID := strings.TrimSpace(talkParts[0])
				toUserID := strings.TrimSpace(talkParts[1])
				// 提取消息内容
				msgContent := strings.TrimSpace(talkParts[2])

				// 发送消息给指定用户
				sendMessageToUser(toUserID, fmt.Sprintf("来自 %s: %s", fromUserID, msgContent))
			} else {
				fmt.Printf("收到无效的聊天消息: %s\n", message)
			}
			continue
		}

		// 处理无效消息
		fmt.Printf("收到用户 %s 的无效消息: %s\n", userID, message)
	}
}

// 删除用户连接
func removeConnection(userID string) {
	mapMutex.Lock()
	delete(userConnMap, userID)
	mapMutex.Unlock()
}
