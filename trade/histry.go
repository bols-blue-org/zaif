package trade

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"net/http"
)

var ZaifKey string
var ZaifSct string

func hamacSha512(data []byte) string {
	hash := hmac.New(sha512.New, []byte(ZaifSct))
	hash.Write([]byte(data))
	//signature := url.QueryEscape(base64.StdEncoding.EncodeToString(hash.Sum(nil)))

	//fmt.Println(signature)
	fmt.Printf("\n%s", data)
	fmt.Printf("\n%x\n", hash.Sum(nil))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func createReqest(body string) *http.Request {
	url := "https://api.zaif.jp/tapi"
	byteBody := []byte(body)
	sigin := hamacSha512(byteBody)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(byteBody))
	headers := http.Header{
		"key":   []string{ZaifKey},
		"sigen": []string{sigin},
	}

	req.Header = headers
	req.Header.Add("key", ZaifKey)
	req.Header.Add("sigin", sigin)
	return req
}

func GetTradeHistory(nonce int) (*http.Response, error) {
	body := fmt.Sprintf("method=trade_history&nonce=%d", nonce)
	req := createReqest(body)
	client := new(http.Client)
	return client.Do(req)
}
func GetInfomation(nonce int) (*http.Response, error) {
	body := fmt.Sprintf("nonce=1501825690.622088&method=get_info")
	req := createReqest(body)
	client := new(http.Client)
	return client.Do(req)
}
