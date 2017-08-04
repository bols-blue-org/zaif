package trade

import (
	"fmt"
	"log"

	"github.com/bols-blue-org/zaif/currency"
)

var tradeMap = MockStatus{map[string]map[string]MockOrder{}}

func CurrencySetTrade(set currency.CurrencySet) {
	if set.AskSimrate() > set.MinWin {
		set.PrintAskTarget()
		price := set.Main.Asks[0].Price + 0.00000001
		tradeMap.UpateOrder(set.Name, "ask", price, set.Unit)
	} else {
		tradeMap.CancelOrder(set.Name, "ask")
	}
	if set.BidSimrate() > set.MinWin {
		set.PrintBidTarget()
		price := set.Main.Bids[0].Price - 0.00000001
		tradeMap.UpateOrder(set.Name, "bid", price, set.Unit)
	} else {
		tradeMap.CancelOrder(set.Name, "bid")
	}
	fmt.Println(tradeMap)
}

type TradeStatus struct {
}

type Trade interface {
	UpateOrder()
	CancelOrder()
	CreateOrder()
	CheckHistry()
}

type MockStatus struct {
	orders map[string]map[string]MockOrder
}

func (ms MockStatus) getOrder(name string, tradeType string) MockOrder {
	if v, ok := ms.orders[name]; !ok {
		ms.orders[name] = map[string]MockOrder{}
		ms.orders[name][tradeType] = MockOrder{
			orderId: 0,
			lastNo:  0,
		}
	} else if _, ok := v[tradeType]; !ok {
		ms.orders[name][tradeType] = MockOrder{
			orderId: 0,
			lastNo:  0,
		}
	}
	return ms.orders[name][tradeType]
}

func (ms *MockStatus) CancelOrder(name string, tradeType string) {
	order := ms.getOrder(name, tradeType)
	if order.orderId != 0 {
		order.lastNo = order.orderId
		order.orderId = 0
		order.price = 0
		order.unit = 0
		ms.orders[name][tradeType] = order
		log.Printf("CancelOrder id %d,%s,%s", order.orderId, name, tradeType)
	}

}
func (ms *MockStatus) CreateOrder(name string, tradeType string, price float64, unit float64) {
	order := ms.getOrder(name, tradeType)
	if order.orderId == 0 {
		order.orderId = order.lastNo + 1
		order.price = price
		order.unit = unit
		ms.orders[name][tradeType] = order
		log.Printf("CreateOrder id %d,%s,%s,price=%.8f,unit=%f", order.orderId, name, tradeType, price, unit)
	}
}
func (ms MockStatus) CheckHistry() {
}
func (ms *MockStatus) UpateOrder(name string, tradeType string, price float64, unit float64) {
	order := ms.getOrder(name, tradeType)
	if order.price != price {
		ms.CancelOrder(name, tradeType)
		ms.CreateOrder(name, tradeType, price, unit)
	}
}

type MockOrder struct {
	orderId int
	lastNo  int
	price   float64
	unit    float64
}
