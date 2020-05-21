package main

import "../../data"
import "../../enums"
import "../../config"

type Collector interface {
    //获取收集器对应的交易所
    getStockExchange() enums.StockExchange
    //获取所有公司信息
    fetchAll(config config.Config) []data.Company
}
