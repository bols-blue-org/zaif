package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bols-blue-org/zaif/currency"
	"github.com/bols-blue-org/zaif/restapi"
	"github.com/bols-blue-org/zaif/stream"

	"golang.org/x/net/websocket"
)

var q (chan currency.CurrencyBoad)

func resevCoinBoad(url string) {
	origin := "http://ws.zaif.jp/"
	var data []byte
	var monaBtc *currency.CurrencyBoad
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	for {
		data = stream.ReadBoad(ws)
		monaBtc, err = currency.NewCurrencyBoad(data)
		if monaBtc != nil {
			q <- *monaBtc
		}
	}
}

const (
	pollingTime = 60
)

func restCallBoad(pairName string) {
	for {
		data, err := restapi.Depth("http://api.zaif.jp/api/1/depth/" + pairName)
		if err == nil {
			boad, err := currency.NewCurrencyBoad(data)
			if boad != nil {
				boad.CurrencyPair = pairName
				q <- *boad
			} else {
				fmt.Printf("%v\n", err)
			}
		}
		time.Sleep(pollingTime * time.Second)
	}
}

var nemSet = currency.CurrencySet{Unit: 100}
var monaSet = currency.CurrencySet{Unit: 100}

func UpdateSet(boad currency.CurrencyBoad) {
	switch boad.CurrencyPair {
	case "mona_jpy":
		monaSet.Sub = &boad
	case "mona_btc":
		monaSet.Main = &boad
		monaSet.PrintSimrate()
	case "btc_jpy":
		monaSet.Btc = &boad
		nemSet.Btc = &boad
	case "xem_jpy":
		nemSet.Sub = &boad
	case "xem_btc":
		nemSet.Main = &boad
		nemSet.PrintSimrate()
	}
}

func main() {
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=xem_btc"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=xem_jpy"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=btc_jpy"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=mona_btc"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=mona_jpy"
	q = make(chan currency.CurrencyBoad, 10)

	go resevCoinBoad("wss://ws.zaif.jp:8888/stream?currency_pair=mona_btc")
	go resevCoinBoad("wss://ws.zaif.jp:8888/stream?currency_pair=mona_jpy")
	go resevCoinBoad("wss://ws.zaif.jp:8888/stream?currency_pair=btc_jpy")
	go resevCoinBoad("wss://ws.zaif.jp:8888/stream?currency_pair=xem_jpy")
	go resevCoinBoad("wss://ws.zaif.jp:8888/stream?currency_pair=xem_btc")
	go restCallBoad("mona_jpy")
	go restCallBoad("mona_btc")
	go restCallBoad("btc_jpy")
	go restCallBoad("xem_jpy")
	go restCallBoad("xem_btc")

	for {
		var boad currency.CurrencyBoad
		select {
		case boad = <-q:
			fmt.Printf("%v\n", boad)
			UpdateSet(boad)
		}

	}

}
