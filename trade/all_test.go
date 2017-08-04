package trade

import (
	"fmt"
	"net/http/httputil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ZaifKey = os.Getenv("ZAIF_KEY")
	ZaifSct = os.Getenv("ZAIF_SCT")
	fmt.Println("key:" + ZaifKey + " sct:" + ZaifSct)
	os.Exit(m.Run())
}

func TestGetTradeHistory(t *testing.T) {
	resp, err := GetTradeHistory(1)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		dumpResp, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("%s", dumpResp)
	}

}

func TestGetInfomation(t *testing.T) {
	resp, err := GetInfomation(2)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		dumpResp, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("%s", dumpResp)
	}

}
func TestOrder(t *testing.T) {
}
