package model

type Position struct {
	Symbol      CurrencyPair
	Side        Direction // 合约类型，Long: 多头，Short: 空头
	Amount      float64   // 持有仓位
	Price       float64   //开仓价格
	ProfitRatio float64   // 盈利比例,正数表示盈利，负数表示亏岁
	Profit      float64   // 盈利

	MarginType     string
	Leverage      float64
	ContractType   string
	ContractId     int64
	ForceLiquPrice float64 //预估爆仓价
	ShortPnlRatio  float64 //空仓收益率
	LongPnlRatio   float64 //多仓收益率
}
