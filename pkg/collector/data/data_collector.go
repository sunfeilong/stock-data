package data

import "github.com/xiaotian/stock/pkg/model"

type Collector interface {
    GetStockExchange() int
    FetchAll(company []model.Company, conf model.StockConfig) []model.Data
}
