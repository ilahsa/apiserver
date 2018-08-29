package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8991")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn.(*net.TCPConn))
	}
}
func handle(server *net.TCPConn) {
	defer server.Close()
	client, err := net.Dial("tcp", "192.168.1.158:9001")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer client.Close()
	//clinet socket 断开，io.CopyBuffer 退出，client 和server 都断开
	go func() {
		defer server.Close()
		defer client.Close()
		buf := make([]byte, 512)
		io.CopyBuffer(server, client, buf)
	}()
	buf := make([]byte, 512)
	io.CopyBuffer(client, server, buf)
}