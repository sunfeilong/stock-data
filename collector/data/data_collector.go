package data

import "../../model"
import "../../config"

type Collector interface {
    GetStockExchange() int
    FetchAll(company []model.Company, conf config.StockConfig) []model.Data
}
