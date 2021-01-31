package main

import (
    "github.com/bittygarden/lilac/io_tool"
    "github.com/xiaotian/stock/pkg/collector"
    "github.com/xiaotian/stock/pkg/config"
    "github.com/xiaotian/stock/pkg/persistent"
    "github.com/xiaotian/stock/pkg/s-logger"
)

var logger = s_logger.New()
var companyFile = persistent.CompanyFilePreserver{}
var dataFile = persistent.DataFilePreserver{}

func main() {
    logger.Infow("检查配置信息")
    path := config.GetDataSaveFilePath()
    if path == "" || io_tool.FileNotExists(path) {
        logger.Errorw("项目运行前检查配置信息，文件保存路径不存在。请在project_config.yml中配置。", "当前配置", path)
        return
    }

    logger.Infow("项目启动")
    companyInfos := collector.CollectCompanyInfo()
    if err := companyFile.Save(companyInfos); err != nil {
        logger.Errorw("保存数据失败", "error", err)
        return
    }

    dataList := collector.CollectData(companyInfos)
    if err := dataFile.Save(dataList); err != nil {
        logger.Errorw("保存数据失败", "error", err)
        return
    }

}
