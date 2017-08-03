package restapi

import (
	"io/ioutil"
	"net/http"
)

//https://api.zaif.jp/api/1/depth/btc_jpy

func Depth(uri string) ([]byte, error) {
	req, _ := http.NewRequest("GET", uri, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}
