package stream

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

func ReadBoad(ws *websocket.Conn) []byte {
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

func PrintBoad(data []byte) {
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
