package collector

import (
    "../config"
    "../model"
    "../s-logger"
    "./company"
)

var logger = s_logger.New()
var collectors map[int]company.Collector = make(map[int]company.Collector)

func init() {
    logger.Infow("初始化收集器容器")
    addToMap(company.SZCompanyCollector{})
}

func addToMap(collector company.Collector) {
    collectors[collector.GetStockExchange()] = collector
}

func CollectCompanyInfo() *[]model.Company {
    logger.Infow("收集公司信息开始")
    tempData := make([]model.Company, 0)
    for key, collector := range collectors {
        logger.Infow("收集公司信息.", "stockCode", key, "collector", collector)
        all := collector.FetchAll(config.GetStockConfig(collector.GetStockExchange()))
        tempData = append(tempData, all...)
    }
    logger.Infow("收集公司信息结束")
    return &tempData
}
