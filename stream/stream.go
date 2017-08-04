package stream

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/websocket"
)

func ReadBoad(ws *websocket.Conn) ([]byte,error) {
	var msg = make([]byte, 1024*4)
	var n int
	var err error
	if n, err = ws.Read(msg); err != nil {
		return nil,err
	}
	var readData = make([]byte, n)
	copy(readData, msg)
	var data []byte
	if n == 4092 {
		if n, err = ws.Read(msg); err != nil {
			return nil,err
		}
		data = append(readData, msg[:n]...)
	}
	if n == 4092 {
		if n, err = ws.Read(msg); err != nil {
			return nil,err
		}
		data = append(readData, msg[:n]...)
	}
	return data,nil
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
