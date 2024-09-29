package server

import "fmt"

// 发送消息给指定用户
func sendMessageToUser(userID, message string) {
	mapMutex.Lock()
	conn, exist := userConnMap[userID]
	mapMutex.Unlock()

	if exist {
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Printf("发送消息给用户 %s 失败: %s\n", userID, err)
			// 删除失效连接
			removeConnection(userID)
		}
	} else {
		fmt.Printf("用户 %s 不在线或不存在\n", userID)
	}
}
