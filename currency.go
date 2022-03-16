package model

import (
	"strings"
)

type Currency struct {
	Symbol string
	Desc   string
}

func (c Currency) String() string {
	return c.Symbol
}

func (c Currency) Eq(c2 Currency) bool {
	return c.Symbol == c2.Symbol
}

var (
	UNKNOWN = Currency{"UNKNOWN", ""}
	CNY     = Currency{"CNY", ""}
	USD     = Currency{"USD", ""}
	USDT    = Currency{"USDT", ""}
	PAX     = Currency{"PAX", "https://www.paxos.com/"}
	USDC    = Currency{"USDC", "https://www.centre.io/"}
	EUR     = Currency{"EUR", ""}
	KRW     = Currency{"KRW", ""}
	JPY     = Currency{"JPY", ""}
	BTC     = Currency{"BTC", "https://bitcoin.org/"}
	XBT     = Currency{"XBT", ""}
	BCC     = Currency{"BCC", ""}
	BCH     = Currency{"BCH", ""}
	BCX     = Currency{"BCX", ""}
	LTC     = Currency{"LTC", ""}
	ETH     = Currency{"ETH", ""}
	ETC     = Currency{"ETC", ""}
	EOS     = Currency{"EOS", ""}
	BTS     = Currency{"BTS", ""}
	QTUM    = Currency{"QTUM", ""}
	SC      = Currency{"SC", ""}
	ANS     = Currency{"ANS", ""}
	ZEC     = Currency{"ZEC", ""}
	DCR     = Currency{"DCR", ""}
	XRP     = Currency{"XRP", ""}
	BTG     = Currency{"BTG", ""}
	BCD     = Currency{"BCD", ""}
	NEO     = Currency{"NEO", ""}
	HSR     = Currency{"HSR", ""}
	BSV     = Currency{"BSV", ""}
	OKB     = Currency{"OKB", "OKB is a global utility token issued by OK Blockchain Foundation"}
	HT      = Currency{"HT", "HuoBi Token"}
	BNB     = Currency{"BNB", "BNB, or Binance Coin, is a cryptocurrency created by Binance."}
	TRX     = Currency{"TRX", ""}
	GBP     = Currency{"GBP", ""}
	XLM     = Currency{"XLM", ""}
	DOT     = Currency{"DOT", ""}
	DASH    = Currency{"DASH", ""}
	CRV     = Currency{"CRV", ""}
	ALGO    = Currency{"ALGO", ""}

	UNKNOWN_PAIR = CurrencyPair{CurrencyA: UNKNOWN, CurrencyB: UNKNOWN}
)

func NewCurrency(symbol, desc string) Currency {
	switch symbol {
	case "cny", "CNY":
		return CNY
	case "usdt", "USDT":
		return USDT
	case "usd", "USD":
		return USD
	case "usdc", "USDC":
		return USDC
	case "pax", "PAX":
		return PAX
	case "jpy", "JPY":
		return JPY
	case "krw", "KRW":
		return KRW
	case "eur", "EUR":
		return EUR
	case "btc", "BTC":
		return BTC
	case "xbt", "XBT":
		return XBT
	case "bch", "BCH":
		return BCH
	case "bcc", "BCC":
		return BCC
	case "ltc", "LTC":
		return LTC
	case "sc", "SC":
		return SC
	case "ans", "ANS":
		return ANS
	case "neo", "NEO":
		return NEO
	case "okb", "OKB":
		return OKB
	case "ht", "HT":
		return HT
	case "bnb", "BNB":
		return BNB
	case "trx", "TRX":
		return TRX
	case "dot", "DOT":
		return DOT
	default:
		return Currency{strings.ToUpper(symbol), desc}
	}
}

