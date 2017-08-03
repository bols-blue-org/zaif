package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//uri := "https://api.zaif.jp/api/1/last_price/btc_jpy"
	uri := "https://api.zaif.jp/api/1/currency_pairs/all"
	req, _ := http.NewRequest("GET", uri, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
