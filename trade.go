package model

import (
	"strings"
	"time"
)

type TradeSide int
type TradeType int
type TradeAct int
type Direction int

const (
	Long  Direction = 1
	Short Direction = 1 << 1
)

const (
	Cancel TradeAct = 1 << 2
	Limit  TradeAct = 1 << 3
	Market TradeAct = 1 << 4
	Stop   TradeAct = 1 << 5
	Open   TradeAct = 1 << 6
	Close  TradeAct = 1 << 7
)

const (
	OpenLong         = TradeType(Open) | TradeType(Long)
	OpenShort        = TradeType(Open) | TradeType(Short)
	CloseLong        = TradeType(Close) | TradeType(Long)
	CloseShort       = TradeType(Close) | TradeType(Short)
	StopLong         = TradeType(Stop) | TradeType(Long)
	StopShort        = TradeType(Stop) | TradeType(Short)
	LimitOpenLong    = TradeType(Open) | TradeType(Long) | TradeType(Limit)
	LimitOpenShort   = TradeType(Open) | TradeType(Short) | TradeType(Limit)
	LimitCloseLong   = TradeType(Close) | TradeType(Long) | TradeType(Limit)
	LimitCloseShort  = TradeType(Close) | TradeType(Short) | TradeType(Limit)
	MarketOpenLong   = TradeType(Open) | TradeType(Long) | TradeType(Market)
	MarketOpenShort  = TradeType(Open) | TradeType(Short) | TradeType(Market)
	MarketCloseLong  = TradeType(Close) | TradeType(Long) | TradeType(Market)
	MarketCloseShort = TradeType(Close) | TradeType(Short) | TradeType(Market)
)

const (
	BUY TradeSide = 1 + iota
	SELL
	BUY_MARKET
	SELL_MARKET
)

func GetDirection(direct string) Direction {
	switch strings.ToUpper(direct) {
	case "SHORT":
		return Short
	}
	return Long
}

func GetTradeSide(ts string) TradeSide {
	switch ts {
	case "BUY":
		return BUY
	case "SELL":
		return SELL
	case "BUY_MARKET":
		return BUY_MARKET
	case "SELL_MARKET":
		return SELL_MARKET
	}
	return TradeSide(0)
}

func (direct Direction) String() string {
	switch direct {
	case Long:
		return "LONG"
	case Short:
		return "SHORT"
	default:
		return "UNKNOWN"
	}
}

func (ts TradeSide) String() string {
	switch ts {
	case 1:
		return "BUY"
	case 2:
		return "SELL"
	case 3:
		return "BUY_MARKET"
	case 4:
		return "SELL_MARKET"
	default:
		return "UNKNOWN"
	}
}

type Trade struct {
	Symbol    CurrencyPair `json:"omitempty"`
	Tid       int64        `json:"tid"`
	Type      TradeSide    `json:"type"`
	Amount    float64      `json:"amount,string"`
	Price     float64      `json:"price,string"`
	Timestamp int64        `json:"date_ms"`
	Date      time.Time
	Side      string
	Remark    string
}

// TradeAction trade action
type TradeAction struct {
	Symbol     CurrencyPair
	Action     TradeType
	Amount     float64
	Price      float64
	Date       time.Time
	OrderID    string
	IsClientID bool
	Timestamp  int64
}

func (a TradeType) IsLong() bool {
	if a&OpenLong == OpenLong || a&CloseShort == CloseShort || a&StopShort == StopShort {
		return true
	}
	return false
}

func (a TradeType) IsOpen() bool {
	if TradeAct(a)&Open == Open {
		return true
	}
	return false
}

func (a TradeType) IsStop() bool {
	if TradeAct(a)&Stop == Stop {
		return true
	}
	return false
}

func (a TradeType) Type() string {
	if TradeAct(a)&Limit == Limit {
		return "L"
	} else if TradeAct(a)&Market == Market {
		return "M"
	}
	return "S"
}

func (a TradeType) Side() (ret string) {
	if TradeAct(a)&Open == Open {
		ret += "O"
	} else if TradeAct(a)&Close == Close {
		ret += "C"
	}

	if Direction(a)&Long == Long {
		ret += "B"
	} else if Direction(a)&Short == Short {
		ret += "S"
	}
	return

}

func (a TradeType) String() (ret string) {
	if TradeAct(a)&Limit == Limit {
		ret += "Limit"
	} else if TradeAct(a)&Market == Market {
		ret += "Market"
	} else if TradeAct(a)&Stop == Stop {
		ret += "Stop"
	}
	if TradeAct(a)&Open == Open {
		ret += "Open"
	} else if TradeAct(a)&Close == Close {
		ret += "Close"
	}

	if Direction(a)&Long == Long {
		ret += "Long"
	} else if Direction(a)&Short == Short {
		ret += "Short"
	}
	return
}
