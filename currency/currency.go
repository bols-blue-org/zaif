package currency

import (
	"encoding/json"
	"fmt"
)

type CurrencyBoad struct {
	Asks []CurrencyLog `json:"asks"`
	Bids []CurrencyLog `json:"bids"`
	Trades  []TradeLog `json:"trades"`
	CurrencyPair string        `json:"currency_pair"`
}

type TradeLog struct {
	Price  float64 `json:"price,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type CurrencyLog struct {
	Price  float64 `json:"price,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

func (cl *CurrencyLog) UnmarshalJSON(value []byte) error {
        data := new([]float64)

        if err := json.Unmarshal(value, data); err != nil {
                return err
        }

    cl.Price = (*data)[0]
    cl.Amount = (*data)[1]

    return nil
}

type CurrencySet struct {
	Main CurrencyBoad
	Sub  CurrencyBoad
	Btc  CurrencyBoad
	Unit float64
}

func (cs *CurrencySet) PrintSimrate(){
	var routePrice = cs.Main.Asks[0].Price * cs.Btc.Asks[0].Price * cs.Unit
	var unitPrice  = cs.Sub.Bids[0].Price * cs.Unit
	fmt.Printf("%.2f - %.2f = win %.2f\n", routePrice, unitPrice, routePrice - unitPrice)
	routePrice = cs.Main.Bids[0].Price * cs.Btc.Bids[0].Price * cs.Unit
	unitPrice  = cs.Sub.Asks[0].Price * cs.Unit
	fmt.Printf("%.2f - %.2f = win %.2f\n", routePrice, unitPrice, routePrice - unitPrice)
}

func NewCurrencyBoad(jsonStr []byte) (*CurrencyBoad, error) {
	jsonBytes := ([]byte)(jsonStr)
	data := new(CurrencyBoad)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}
