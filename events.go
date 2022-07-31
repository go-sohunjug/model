package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-sohunjug/indicator"
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

var EventTypes = map[string]reflect.Type{
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

type Engine interface {
	OpenLong(action *TradeAction) *Order
	CloseLong(action *TradeAction) *Order
	OpenShort(action *TradeAction) *Order
	CloseShort(action *TradeAction) *Order
	StopLong(action *TradeAction) *Order
	StopShort(action *TradeAction) *Order
	GetOrder(action *TradeAction) *Order
	CancelAllOrder(action *TradeAction)
	CancelOrder(action *TradeAction) bool
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
	// SetTag(key, value string)
	Filter(key, value string) bool
	Check(key, value string) bool
}

type Runner interface {
	Filter(name, key string) bool
	Param() (params ParamData)
	Init(engine Engine, params map[string]interface{})
	OnTick(tick *Ticker)
	OnCandle(candle *Candle)
	OnPosition(position *Position)
	OnTrade(trade *Trade)
	OnTrades(trade *Trade)
	OnDepth(depth *Depth)
	OnAccount(account *Account)
	OnOrder(order *Order)
	// OnEvent(e Event) (err error)
	UpdateParams(params map[string]interface{})
	GetName() string
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
	Symbol *CurrencyPair
}
