package model

import "time"

type DepthRecord struct {
	Price  float64
	Amount float64
	Total  float64
}

type Depth Orderbook
type DepthRecords []DepthRecord

type Orderbook struct {
	Symbol    CurrencyPair
	Asks      DepthRecords
	Bids      DepthRecords
	Timestamp int64
	Date      time.Time
}

func (dr DepthRecords) Len() int {
	return len(dr)
}

func (dr DepthRecords) Swap(i, j int) {
	dr[i], dr[j] = dr[j], dr[i]
}

func (dr DepthRecords) Less(i, j int) bool {
	return dr[i].Price < dr[j].Price
}


