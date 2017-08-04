package trade

import "testing"

func TestMockTrade(t *testing.T) {
	tmp := tradeMap
	tmp.CreateOrder("mona", "ask", 0.1, 0.1)
	tmp.CreateOrder("mona", "ask", 0.1, 0.1)
	tmp.CreateOrder("mona", "ask", 0.1, 0.1)
	tmp.CreateOrder("mona", "ask", 0.1, 0.1)
	if tmp.getOrder("mona", "ask").orderId != 1 {
		t.Errorf("error %v", tmp)
	}
	tmp.CancelOrder("mona", "ask")
	tmp.CancelOrder("mona", "ask")
	tmp.CancelOrder("mona", "ask")
	tmp.CancelOrder("mona", "ask")
	tmp.CancelOrder("mona", "ask")
	if tmp.getOrder("mona", "ask").orderId != 0 {
		t.Errorf("error %v", tmp)
	}
	if tmp.getOrder("mona", "ask").lastNo != 1 {
		t.Errorf("error %v", tmp)
	}
	tmp.CreateOrder("mona", "ask", 0.1, 0.1)
	if tmp.getOrder("mona", "ask").orderId != 2 {
		t.Errorf("error %v", tmp)
	}
	if tmp.getOrder("mona", "ask").lastNo != 1 {
		t.Errorf("error %v", tmp)
	}
	tmp.CancelOrder("mona", "ask")
	if tmp.getOrder("mona", "ask").orderId != 0 {
		t.Errorf("error %v", tmp)
	}
	if tmp.getOrder("mona", "ask").lastNo != 2 {
		t.Errorf("error %v", tmp)
	}

}
