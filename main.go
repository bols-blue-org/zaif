package main

import (
	"fmt"
	"log"

	"github.com/bols-blue-org/zaif/currency"
	"github.com/bols-blue-org/zaif/stream"

	"golang.org/x/net/websocket"
)

func main() {
	origin := "http://ws.zaif.jp/"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=xem_btc"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=xem_jpy"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=btc_jpy"
	url := "wss://ws.zaif.jp:8888/stream?currency_pair=mona_btc"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=mona_jpy"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	var monaBtc *currency.CurrencyBoad
	for {
		data = stream.ReadBoad(ws)
		monaBtc, err = currency.NewCurrencyBoad(data)
		fmt.Printf("%v\n", monaBtc)
	}
}
