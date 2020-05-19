package company

import (
	"../../data"
	"../../enums"
)

//深圳交易所上市公司信息收集器
type SZCompanyCollector struct {
}

func (sz SZCompanyCollector) getStockExchange() enums.StockExchange {
	return enums.SZ
}

func (sz SZCompanyCollector) fetchAll() []data.Company {
	return nil
}
