package main

import (
    "./collector"
    "./persistent"
    "./s-logger"
)

var logger = s_logger.New()
var companyFile = persistent.CompanyFilePreserver{}
var dataFile = persistent.DataFilePreserver{}

func main() {
    logger.Infow("项目启动")
    logger.Infow("收集公司信息触发执行")
    companyInfos := collector.CollectCompanyInfo()
    if err := companyFile.Save(companyInfos); err != nil {
        logger.Errorw("保存数据失败", "error", err)
    }

    dataList := collector.CollectData(companyInfos)
    if err := dataFile.Save(dataList); err != nil {
        logger.Errorw("保存数据失败", "error", err)
    }
    logger.Infow("项目运行结束")
}
