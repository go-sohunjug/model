package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-sohunjug/quant/v1/indicator"
)

// Events
const (
	EventCandleParam    = "event.candle_param"
	EventCandle         = "event.candle"
	EventTicker         = "event.ticker"
	EventOrder          = "event.order"
	EventOrderCancelAll = "event.order_cancel_all"
	EventOrderCancel    = "event.order_cancel"
	EventTrades         = "event.trades"
	// own trades
	EventTradeAction = "event.trade_action"
	EventTrade       = "event.trade"
	EventPosition    = "event.position"
	EventCurPosition = "event.cur_position" // position of current script
	EventRiskLimit   = "event.risk_limit"
	EventDepth       = "event.depth"
	// all trades in the markets

	EventAccount = "event.balance"
	EventAction  = "event.action"

	EventWatch = "event.watch"

	EventNotify = "event.notify"
)

var (
	EventTypes = map[string]reflect.Type{
		EventCandleParam: reflect.TypeOf(CandleParam{}),
		EventCandle:      reflect.TypeOf(Candle{}),
		EventTicker:      reflect.TypeOf(Ticker{}),
		EventOrder:       reflect.TypeOf(Order{}),
		EventOrderCancel: reflect.TypeOf(TradeAction{}),
		// EventOrderCancelAll     = "order_cancel_all"
		EventTrade:       reflect.TypeOf(Trade{}),
		EventTrades:      reflect.TypeOf(Trade{}),
		EventTradeAction: reflect.TypeOf(TradeAction{}),
		EventPosition:    reflect.TypeOf(Position{}),
		// EventCurPosition        = "cur_position" // position of current script
		// EventRiskLimit          = "risk_limit"
		EventDepth:   reflect.TypeOf(Depth{}),
		EventAccount: reflect.TypeOf(Account{}),
		EventAction:  reflect.TypeOf(EngineAction{}),

		EventNotify: reflect.TypeOf(NotifyEvent{}),
	}
)

type Engine interface {
	OpenLong(symbol CurrencyPair, price, amount float64)
	CloseLong(symbol CurrencyPair, price, amount float64)
	OpenShort(symbol CurrencyPair, price, amount float64)
	CloseShort(symbol CurrencyPair, price, amount float64)
	StopLong(symbol CurrencyPair, price, amount float64)
	StopShort(symbol CurrencyPair, price, amount float64)
	GetOrder(symbol, orderid string)
	CancelAllOrder()
	CancelOrder(symbol CurrencyPair, orderId string, isClientId bool)
	AddIndicator(name string, params ...int) (ind indicator.CommonIndicator)
	Log(v ...interface{})
	Logf(f string, v ...interface{})
	Watch(watchType string)
	SendNotify(content, contentType string)

	Start()
	Stop()
	SaveParams()
	// call for goscript
	AddTimer(second int64, timer func())
	OnCandle(candle *Candle)
	SetTag(key, value string)
	Filter(key, value string) bool
}

// CandleParam get candle param
type CandleParam struct {
	Start    time.Time
	End      time.Time
	Exchange string
	BinSize  string
	Symbol   string
}

// NotifyEvent event to send notify
type NotifyEvent struct {
	Type    string // text,markdown
	Content string
}

// RiskLimit risk limit
type RiskLimit struct {
	Code         string  // symbol info, empty = global
	Lever        float64 // lever
	MaxLostRatio float64 // max lose ratio
}

// Key key of r
func (r RiskLimit) Key() string {
	return fmt.Sprintf("%s-%.2f", r.Code, r.Lever)
}

// WatchParam add watch event param
type WatchParam struct {
	Type  string
	Param map[string]interface{}
}

type EngineAction struct {
	Action string
	Name   string
	Symbol string
}
