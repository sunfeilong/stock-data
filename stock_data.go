package main

import (
    "./collector"
    "./persistent"
    "./s-logger"
)

var logger = s_logger.New()
var preserver = persistent.CompanyFilePreserver{}

func main() {
    logger.Infow("项目启动")
    logger.Infow("收集公司信息触发执行")
    companyInfos := collector.CollectCompanyInfo()
    if err := preserver.Save(companyInfos); err != nil {
        logger.Errorw("保存数据失败", "error", err)
    }
    logger.Infow("项目运行结束")
}
