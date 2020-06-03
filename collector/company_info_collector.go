package collector

import (
    "../config"
    "../enums"
    "../model"
    "../persistent"
    "../s-logger"
    "./company"
)

var logger = s_logger.New()
var collectors map[int]company.Collector = make(map[int]company.Collector)
var preserver = persistent.CompanyFilePreserver{}

func init() {
    logger.Infow("初始化收集器容器")
    addToMap(company.SZCompanyCollector{})
}

func addToMap(collector company.Collector) {
    collectors[collector.GetStockExchange()] = collector
}

func CollectCompanyInfo() *[]model.Company {
    logger.Infow("收集公司信息开始")
    plateEnums := enums.GetAll()
    tempData := make([]model.Company, 0)
    for _, p := range plateEnums {
        collector := collectors[p.StockExchange]
        if nil != collector {
            all := collector.FetchAll(config.GetStockConfig(collector.GetStockExchange()))
            tempData = append(tempData, all...)
        } else {
            logger.Infow("收集公司信息,没有找到对应的收集器")
        }
    }
    if err := preserver.Save(tempData); err != nil {
        logger.Errorw("收集公司信息,保存数据失败", "error", err)
    }
    logger.Infow("收集公司信息结束")
    return &tempData
}
