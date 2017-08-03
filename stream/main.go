package main

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

func readBoad(ws *websocket.Conn) []byte {
	var msg = make([]byte, 1024*4)
	var n int
	var err error
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	var readData = make([]byte, n)
	copy(readData, msg)
	var data []byte
	if n == 4092 {
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		data = append(readData, msg[:n]...)
	}
	if n == 4092 {
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		data = append(readData, msg[:n]...)
	}
	return data
}

func printBoad(data []byte) {
	var query map[string]interface{}
	json.Unmarshal(data, &query)
	if query["bids"] != nil {
		var bids []interface{}
		bids = query["bids"].([]interface{})
		fmt.Printf("%v\n", bids[:3])

	}
	if query["asks"] != nil {
		var asks []interface{}
		asks = query["asks"].([]interface{})
		fmt.Printf("%v\n", asks[:3])
	}
}

func main() {
	origin := "http://ws.zaif.jp/"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=xem_btc"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=xem_jpy"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=btc_jpy"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=mona_btc"
	url := "wss://ws.zaif.jp:8888/stream?currency_pair=mona_jpy"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	/*
		var message string
		websocket.Message.Receive(ws, &message)
		fmt.Printf("%s\n", message)
	*/
	data := readBoad(ws)
	fmt.Printf("%s",data)
	//printBoad(readBoad(ws))
	//printBoad(readBoad(ws))
	//printBoad(readBoad(ws))
	//printBoad(readBoad(ws))
	//printBoad(readBoad(ws))
	//printBoad(readBoad(ws))
}
