package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

var subJsonStr = `
{
          "name": "xem_jpy",
	  "asks": [[24.3999, 9902.0], [24.4999, 7856.0], [24.5, 351.0], [24.56, 15.0], [24.57, 15.0], [24.66, 15.0], [24.74, 6897.3], [24.76, 15.0], [24.7999, 7269.0], [24.8, 13000.0], [24.8597, 3.0], [24.86, 30.0], [24.8871, 500.0], [24.8998, 1741.0], [24.8999, 57308.0], [24.99, 4500.0], [24.9998, 1500.0], [24.9999, 4036.0], [25.0, 112044.0], [25.0799, 61403.0], [25.08, 500.0], [25.1, 1856.0], [25.15, 2273.0], [25.1691, 64484.0], [25.1697, 12813.0], [25.1699, 5000.0], [25.17, 15.0], [25.189, 120.0], [25.19, 630.0], [25.1999, 400.0], [25.2, 57779.0], [25.22, 15.0], [25.25, 3905.0], [25.27, 15.0], [25.2999, 430.0], [25.3, 9730.0], [25.31, 15.0], [25.33, 30150.0], [25.4, 2000.0], [25.468, 600.0], [25.475, 10000.0], [25.48, 1776.0], [25.4843, 100.0], [25.5, 76728.0], [25.5496, 100.0], [25.5774, 100.0], [25.586, 1.0], [25.6, 5.0], [25.6523, 2248.0], [25.7, 3160.0], [25.7133, 100.0], [25.7888, 3000.0], [25.8, 1550.0], [25.8951, 100.0], [25.9, 19925.0], [25.92, 15981.0], [25.99, 100.0], [25.9995, 1000.0], [25.9999, 20170.0], [26.0, 197915.8], [26.01, 15.0], [26.02, 15.0], [26.03, 15.0], [26.07, 15.0], [26.09, 15.0], [26.11, 15.0], [26.12, 15.0], [26.15, 3905.0], [26.16, 15.0], [26.18, 30.0], [26.19, 15.0], [26.2, 5.0], [26.21, 40512.0], [26.3, 6696.0], [26.31, 189.0], [26.32, 15.0], [26.3333, 33.3], [26.39, 15.0], [26.4, 11942.0], [26.49, 15.0], [26.5, 16015.0], [26.52, 15.0], [26.54, 15.0], [26.55, 15.0], [26.6, 20.0], [26.61, 15.0], [26.62, 15.0], [26.72, 15.0], [26.74, 15.0], [26.77, 17550.0], [26.8, 30.0], [26.82, 15.0], [26.84, 15.0], [26.85, 200.0], [26.86, 15.0], [26.87, 3500.0], [26.879, 6195.0], [26.88, 15.0], [26.8999, 3000.0], [26.9, 755.0], [26.92, 15.0], [26.99, 355.0], [26.9995, 1000.0], [26.9999, 16615.0], [27.0, 247806.0], [27.01, 15.0], [27.02, 15.0], [27.03, 15.0], [27.04, 15.0], [27.09, 1144.8], [27.1, 30000.0], [27.2, 5.0], [27.22, 15.0], [27.26, 15.0], [27.3, 20.0], [27.32, 15.0], [27.38, 723.0], [27.4, 10705.0], [27.45, 753.0], [27.48, 1381.0], [27.487, 10000.0], [27.495, 50240.0], [27.5, 28142.0], [27.51, 15.0], [27.56, 17131.0], [27.595, 1300.0], [27.5997, 250.0], [27.6, 5.0], [27.61, 100000.0], [27.78, 4946.0], [27.88, 4350.0], [27.89, 30000.0], [27.9, 2972.0], [28.0, 57876.8], [28.11, 50000.0], [28.115, 800.0], [28.15, 1000.0], [28.2, 5.0], [28.24, 2129.6], [28.42, 765.0], [28.5, 23010.0], [28.6, 5.0], [28.8, 32217.1], [28.88, 2000.0], [28.9, 2576.0], [28.9999, 500.0], [29.0, 130210.0], [29.1, 150.0], [29.2, 38888.0], [29.4, 5398.0]],
	  "bids": [[24.3006, 2000.0], [24.3005, 30171.0], [24.3003, 258.0], [24.3002, 15858.0], [24.3001, 19417.0], [24.3, 785.9], [24.2557, 17066.0], [24.2556, 636.0], [24.2555, 15301.0], [24.25, 1.0], [24.234, 11246.0], [24.224, 882.0], [24.22, 12502.0], [24.2001, 100.0], [24.2, 46462.0], [24.15, 37219.0], [24.122, 2594.0], [24.11, 5745.0], [24.1, 14482.0], [24.09, 10633.0], [24.0101, 7536.9], [24.006, 20000.0], [24.003, 1784.0], [24.001, 1000.0], [24.0005, 1000.0], [24.0002, 13191.0], [24.0001, 15237.0], [24.0, 69911.0], [23.9998, 690174.0], [23.95, 10000.0], [23.9, 18044.0], [23.886, 800.0], [23.8141, 100.0], [23.81, 40.0], [23.8, 6804.0], [23.7, 10058.0], [23.69, 15.0], [23.68, 20000.0], [23.67, 15.0], [23.6683, 50.0], [23.6606, 690174.0], [23.65, 3042.0], [23.63, 5015.0], [23.6001, 3.0], [23.6, 1469.0], [23.5601, 71.0], [23.56, 2000.0], [23.5201, 500.0], [23.5002, 100.0], [23.5001, 7000.0], [23.5, 15065.0], [23.46, 2561.0], [23.45, 15.0], [23.43, 15.0], [23.42, 15.0], [23.4, 2000.0], [23.39, 15.0], [23.38, 500.0], [23.35, 15.0], [23.34, 2400.0], [23.33, 500.0], [23.32, 15.0], [23.3101, 5.0], [23.31, 117.0], [23.3, 1015.0], [23.26, 5000.0], [23.25, 15.0], [23.21, 2500.0], [23.2, 300.0], [23.19, 214.9], [23.18, 500.0], [23.11, 15.0], [23.1, 54036.0], [23.09, 15.0], [23.07, 500.0], [23.05, 15.0], [23.001, 2300.0], [23.0, 80001.0], [22.9, 15.0], [22.87, 2500.0], [22.86, 1030.0], [22.85, 9032.0], [22.82, 15.0], [22.8, 1753.0], [22.7001, 1000.0], [22.7, 1000.0], [22.65, 15.0], [22.6401, 5896.0], [22.61, 2000.0], [22.6, 10216.0], [22.55, 32466.0], [22.51, 5000.0], [22.5, 38177.0], [22.48, 16.0], [22.47, 15.0], [22.46, 15.0], [22.45, 91.0], [22.4, 338.0], [22.39, 15.0], [22.35, 17.0], [22.3, 15.0], [22.28, 15.0], [22.2, 570.0], [22.13, 15.0], [22.12, 15.0], [22.1, 55271.0], [22.05, 15.0], [22.001, 133.0], [22.0005, 502.0], [22.0, 60783.0], [21.999, 10000.0], [21.99, 15.0], [21.97, 0.1], [21.93, 15.0], [21.92, 15.0], [21.91, 15.0], [21.9, 15.0], [21.89, 15.0], [21.8811, 1000.0], [21.87, 15.0], [21.85, 15.0], [21.82, 15.0], [21.8, 21275.0], [21.78, 15.0], [21.75, 15.0], [21.72, 20257.0], [21.71, 15.0], [21.7, 8015.0], [21.69, 15.0], [21.68, 15.0], [21.67, 15.0], [21.66, 15.0], [21.64, 15.0], [21.63, 15.0], [21.62, 417.0], [21.59, 15.0], [21.58, 15.0], [21.5501, 10000.0], [21.55, 1800.0], [21.53, 15.0], [21.5005, 1000.0], [21.5, 7144.0], [21.46, 500.0], [21.45, 15.0], [21.44, 15.0], [21.42, 15.0], [21.41, 15.0], [21.4, 15.0], [21.38, 14.0], [21.37, 15.0]],
          "log":[
                {"price":22.8, "amount":100},
                {"price":22.8, "amount":100},
                {"price":22.8, "amount":100},
                {"price":22.8, "amount":100},
                {"price":22.8, "amount":100}
          ]
        }`

const bufSize = 1024 * 16

func readFile(path string) []byte {
    file, err := os.Open(path)
    if err != nil {
            fmt.Println(err)
    }
    defer file.Close()

    buf := make([]byte, bufSize)
    for {
        n, err := file.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            fmt.Println(err)
            break
        }
        return buf[:n]
    }
    return buf
}

func TestMain(m *testing.M) {
	jsonStr := `{
	  "name": "日本",
	  "asks": [
	  ],
	  "bids":[
	  ],
	  "log":[
		{"price":0.02, "amount":100},
		{"price":0.02, "amount":100},
		{"price":0.02, "amount":100},
		{"price":0.02, "amount":100},
		{"price":0.02, "amount":100}
	  ]
	}`
	jsonBytes := ([]byte)(jsonStr)
	data := new(CurrencyBoad)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	fmt.Printf("%v\n", data)
	os.Exit(m.Run())
}

func TestCrrencySet(t *testing.T) {
	var jsonString = readFile("./data/xem_btc.json")
	nemBtc, err := NewCurrencyBoad(jsonString)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	} else {
		fmt.Printf("%v\n", nemBtc)
	}

	nemJpy, err := NewCurrencyBoad(([]byte)(subJsonStr))
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	} else {
		fmt.Printf("%v\n", nemJpy)
	}

	jsonString = readFile("./data/btc_jpy.json")
	btcJpy, err := NewCurrencyBoad(jsonString)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	} else {
		fmt.Printf("%v\n", btcJpy)
	}
	set := CurrencySet{Main:*nemBtc,Sub:*nemJpy,Btc:*btcJpy,Unit:100}
	set.PrintSimrate()
}
