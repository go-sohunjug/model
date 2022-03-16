package model

import "time"

type Ticker struct {
	Symbol    CurrencyPair `json:"omitempty"`
	Open      float64      `json:"open,string"`
	Last      float64      `json:"last,string"`
	Bid       float64      `json:"buy,string"`
	Ask       float64      `json:"sell,string"`
	High      float64      `json:"high,string"`
	Low       float64      `json:"low,string"`
	BaseVol   float64      `json:"baseVol,string"`
	QuoteVol  float64      `json:"quoteVol,string"`
	LastVol   float64      `json:"lastVol,string"`
	Timestamp int64        `json:"date"` // 单位:ms
	Date      time.Time    `json:"datetime,omitempty"`
}
