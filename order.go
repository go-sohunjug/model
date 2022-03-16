package model

import "time"

var (
	OrderStatusFilled   = "FILLED"
	OrderStatusCanceled = "CANCELED"
)

type TradeStatus int

func (ts TradeStatus) String() string {
	return tradeStatusSymbol[ts]
}

func GetTradeStatus(ts string) TradeStatus {
	switch ts {
	case tradeStatusSymbol[0]:
		return TradeStatus(0)
	case tradeStatusSymbol[1]:
		return TradeStatus(1)
	case tradeStatusSymbol[2]:
		return TradeStatus(2)
	case tradeStatusSymbol[3]:
		return TradeStatus(3)
	case tradeStatusSymbol[4]:
		return TradeStatus(4)
	case tradeStatusSymbol[5]:
		return TradeStatus(5)
	case tradeStatusSymbol[6]:
		return TradeStatus(6)
	case "FILLED":
		return TradeStatus(2)
	case "EXPIRED":
		return TradeStatus(3)
	case "PARTIALLY_FILLED":
		return TradeStatus(1)
	}

	return TradeStatus(0)
}

var tradeStatusSymbol = [...]string{"UNFINISH", "PART_FINISH", "FINISH", "CANCELED", "REJECT", "CANCEL_ING", "FAIL"}

const (
	ORDER_UNFINISH TradeStatus = iota
	ORDER_PART_FINISH
	ORDER_FINISH
	ORDER_CANCEL
	ORDER_REJECT
	ORDER_CANCEL_ING
	ORDER_FAIL
)

const (
	OPEN_BUY   = "OB" //开多
	OPEN_SELL  = "OS" //开空
	CLOSE_BUY  = "CB" //平多
	CLOSE_SELL = "CS" //平空
)

type Order struct {
	ClientOrderID string
	OrderID       string
	Symbol        CurrencyPair
	Price         float64
	Amount        float64
	AvgPrice      float64
	DealAmount    float64
	Fee           float64
	Status        TradeStatus
	Side          TradeSide
	Offset        string
	OrderType     int   //0:default,1:maker,2:fok,3:ioc
	OrderTime     int64 // create  timestamp
	FinishedTime  int64 //finished timestamp
	//策略委托单
	LeverRate    float64 //倍数
	TriggerPrice float64
	AlgoType     int //1:限价 2:市场价；触发价格类型，默认是限价；为市场价时，委托价格不必填；

	Timestamp int64
	Date      time.Time
}
