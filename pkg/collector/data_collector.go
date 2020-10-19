package collector

import (
    "github.com/xiaotian/stock/pkg/collector/data"
    "github.com/xiaotian/stock/pkg/config"
    "github.com/xiaotian/stock/pkg/model"
)

var dataCollectors map[int]data.Collector = make(map[int]data.Collector)

func init() {
    logger.Infow("初始化数据收集器容器")
    addDataCollectorToMap(data.SZDataCollector{})
    addDataCollectorToMap(data.SHDataCollector{})
    addDataCollectorToMap(data.HKDataCollector{})
}

func addDataCollectorToMap(collector data.Collector) {
    dataCollectors[collector.GetStockExchange()] = collector
}

func CollectData(company []model.Company) []model.Data {
    logger.Infow("收集公司数据开始")
    tempData := make([]model.Data, 0)
    for key, collector := range dataCollectors {
        logger.Infow("收集公司数据.", "stockCode", key, "collector", collector)
        all := collector.FetchAll(company, config.GetStockConfig(collector.GetStockExchange()))
        tempData = append(tempData, all...)
    }
    logger.Infow("收集公司数据结束")
    return tempData
}
