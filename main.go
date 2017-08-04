package main

import (
	"fmt"
	"log"
	"log/syslog"
	"time"

	"github.com/bols-blue-org/zaif/currency"
	"github.com/bols-blue-org/zaif/restapi"
	"github.com/bols-blue-org/zaif/stream"
	"github.com/bols-blue-org/zaif/trade"

	"golang.org/x/net/websocket"
)

var q (chan currency.CurrencyBoad)

func resevCoinBoad(url string) {
	origin := "http://ws.zaif.jp/"
	var data []byte
	var monaBtc *currency.CurrencyBoad
restart:
	log.Println(url)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Println(err)
		time.Sleep(1 * time.Second)
		goto restart
	}
	for {
		data, err = stream.ReadBoad(ws)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			goto restart
		} else {
			monaBtc, err = currency.NewCurrencyBoad(data)
			if monaBtc != nil {
				q <- *monaBtc
			}
		}
	}
}

const (
	pollingTime = 10
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

var nemSet = currency.CurrencySet{Unit: 100, Name: "nem", MinWin: 20}
var monaSet = currency.CurrencySet{Unit: 100, Name: "mona", MinWin: 50}

func UpdateSet(boad currency.CurrencyBoad) {
	switch boad.CurrencyPair {
	case "mona_jpy":
		monaSet.Sub = boad
	case "mona_btc":
		monaSet.Main = boad
		mainChan <- monaSet
	case "btc_jpy":
		monaSet.Btc = boad
		nemSet.Btc = boad
	case "xem_jpy":
		nemSet.Sub = boad
	case "xem_btc":
		nemSet.Main = boad
		mainChan <- nemSet

	}
}

var mainChan (chan currency.CurrencySet)

func CurrencyTrade() {
	for {
		var set currency.CurrencySet
		select {
		case set = <-mainChan:
			trade.CurrencySetTrade(set)
		}

	}
}

func main() {
	logger, err := syslog.New(syslog.LOG_NOTICE|syslog.LOG_USER, "zaif-daemon")
	if err != nil {
		panic(err)
	}
	log.SetOutput(logger)
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=xem_btc"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=xem_jpy"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=btc_jpy"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=mona_btc"
	//url := "wss://ws.zaif.jp:8888/stream?currency_pair=mona_jpy"
	q = make(chan currency.CurrencyBoad, 10)
	mainChan = make(chan currency.CurrencySet, 10)

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
	go CurrencyTrade()

	for {
		var boad currency.CurrencyBoad
		select {
		case boad = <-q:
			if boad.CurrencyPair != "btc_jpy" {
				fmt.Printf("%v\n", boad)
			}
			UpdateSet(boad)
		}

	}

}
