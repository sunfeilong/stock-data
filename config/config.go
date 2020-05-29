package config

import (
    "github.com/spf13/viper"
)

//配置文件名字
const (
    defaultConfigFileName string = "project_config"
    logConfigFileName     string = "log"
)

var projectViper *viper.Viper
var stockConfigMap = make(map[int]StockConfig)

func init() {
    //项目配置文件
    projectViper = viper.New()
    projectViper.SetConfigType("yaml")
    projectViper.AddConfigPath(".")
    projectViper.AddConfigPath("../")
    projectViper.AddConfigPath("../../")
    projectViper.SetConfigName(defaultConfigFileName)
    if err := projectViper.ReadInConfig(); nil != err {
        panic(err)
    }
    //股票信息配置
    var P *StockConfigs
    if err := projectViper.UnmarshalKey("stock", &P); nil != err {
        panic(err)
    }
    for _, stockConfig := range P.Configs {
        stockConfigMap[stockConfig.StockExchangeCode] = stockConfig
    }
}

func GetStockConfig(s int) StockConfig {
    return stockConfigMap[int(s)]
}

func GetLogConfig() {

}
