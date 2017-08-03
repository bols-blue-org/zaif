package restapi

import (
	"fmt"
	"testing"

	"github.com/bols-blue-org/zaif/currency"
)

func TestGetDepth(t *testing.T) {
	data, err := Depth("http://api.zaif.jp/api/1/depth/mona_jpy")
	if err == nil {
		fmt.Printf("%s", data)
		boad, err := currency.NewCurrencyBoad(data)
		if boad != nil {
			fmt.Printf("%v\n", boad)
		} else {
			fmt.Printf("%v\n", err)
		}

	}
}
