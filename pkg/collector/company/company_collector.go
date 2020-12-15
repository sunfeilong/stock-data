package company

import (
    "github.com/xiaotian/stock/pkg/model"
)

type Collector interface {
    //获取收集器对应的交易所
    GetStockExchange() int
    //获取所有公司信息
    FetchAll(config model.StockConfig) []model.Company
}
