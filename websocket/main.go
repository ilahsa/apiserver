package main

import (
	"golang.org/x/net/websocket"
	"log"
	"fmt"
	"time"
	"strconv"
)

func main() {
	ch1 := make(chan bool)
	roomId:=strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	registerInfo1:=fmt.Sprintf(`{"cmd":"register","roomid":"%s","clientid":"aaaaaaaaaaaa"}`,roomId)
	registerInfo2:=fmt.Sprintf(`{"cmd":"register","roomid":"%s","clientid":"bbbbbbbbbbbb"}`,roomId)
	ws1, err := register(registerInfo1)
	go func() {
		var msg = make([]byte, 1024)
		var n int
		if n, err = ws1.Read(msg); err != nil {
			log.Fatal(err)
		}
		t := time.Now()
		timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
		fmt.Println(timestamp)
		fmt.Printf("Received: %s.\n", msg[:n])
		<-ch1
	}()
	ws2, err := register(registerInfo2)
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	fmt.Println(timestamp)
	ws2.Write([]byte(`{"cmd":"send","msg":"{\"type\": \"candidate\";\"tag\":\"websocket_test_req\"}"}`))

	ch1 <- true

}

func register(registerInfo string) (*websocket.Conn, error) {
	origin := "http://localhost/"
	url := "ws://sy.big.chumob.com/ws"
	//url := "ws://192.168.0.128:8181/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if _, err := ws.Write([]byte(registerInfo)); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return ws, nil
}
