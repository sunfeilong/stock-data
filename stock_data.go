package main

import (
    "github.com/xiaotian/stock/pkg/collector"
    "github.com/xiaotian/stock/pkg/persistent"
    "github.com/xiaotian/stock/pkg/s-logger"
)

var logger = s_logger.New()
var companyFile = persistent.CompanyFilePreserver{}
var dataFile = persistent.DataFilePreserver{}

func main() {
    logger.Infow("项目启动")
    companyInfos := collector.CollectCompanyInfo()
    if err := companyFile.Save(companyInfos); err != nil {
        logger.Errorw("保存数据失败", "error", err)
    }

    dataList := collector.CollectData(companyInfos)
    if err := dataFile.Save(dataList); err != nil {
        logger.Errorw("保存数据失败", "error", err)
    }

}
