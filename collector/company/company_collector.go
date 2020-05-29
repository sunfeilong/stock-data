package company

import "../../data"
import "../../config"

type Collector interface {
    //获取收集器对应的交易所
    getStockExchange() int
    //获取所有公司信息
    fetchAll(config config.StockConfig) []data.Company
}
