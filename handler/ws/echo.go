package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)




var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func EchoMessage(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	for {
		// 读取客户端的消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// 把消息打印到标准输出
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// 把消息写回客户端，完成回音
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
