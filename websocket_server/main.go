package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

func worker(ch chan int, ws *websocket.Conn) {
	var err error
	for {
		select {
		case <-ch:
			return //收到信号就退出线程
		default:
			msg := "server to client " + time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("send to client:" + msg)
			if err = websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("send failed:", err)
				break
			}
			time.Sleep(time.Duration(10) * time.Second)
		}
	}
}

func Echo(ws *websocket.Conn) {
	//这里是发送消息
	var err error
	ch := make(chan int)
	for {
		var reply string
		//websocket接受信息
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}
		msg1 := "received " + reply
		if reply == "start_record" {
			go worker(ch, ws)
			websocket.Message.Send(ws, "success")
			continue
		}
		if reply == "stop_record" {
			ch <- 1 //发送退出线程的命令
			websocket.Message.Send(ws, "success")
			continue
		}
		websocket.Message.Send(ws, msg1)
		fmt.Println("reveived from client: " + reply)
	}
}
func main() {
	//接受websocket的路由地址
	http.Handle("/websocket", websocket.Handler(Echo))
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
