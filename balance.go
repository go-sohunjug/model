package model

type Account struct {
	Currency         Currency
	Available        float64
	Frozen           float64
	Total            float64
	MaintMargin      float64
	ProfitUnreal float64
}

type Balance map[Currency]Account

type BalanceInfo struct {
	Exchange string
	Method   string
	Asset    float64 //总资产
	NetAsset float64 //净资产
	Detail   Balance
}
