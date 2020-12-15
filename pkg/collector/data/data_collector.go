package data

import "github.com/xiaotian/stock/pkg/model"
import "github.com/xiaotian/stock/pkg/config"

type Collector interface {
    GetStockExchange() int
    FetchAll(company []model.Company, conf config.StockConfig) []model.Data
}
