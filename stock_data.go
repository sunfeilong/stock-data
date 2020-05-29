package main

import (
    "./collector"
    "./s-logger"
)

var logger = s_logger.New()

func main() {
    logger.Infow("项目启动")
    collector.CollectCompanyInfo()
    logger.Infow("项目运行结束")
}
