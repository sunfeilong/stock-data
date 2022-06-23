package config

import (
	"github.com/spf13/viper"
	"github.com/xiaotian/stock/model"
)

//配置文件名字
const (
	defaultConfigFileName string = "project_config"
)

var projectViper *viper.Viper
var stockConfigMap = make(map[int]model.StockConfig)
var dataSaveFilePath string
var skipIfNoData bool

func init() {
	//项目配置文件
	projectViper = viper.New()
	projectViper.SetConfigType("yaml")
	projectViper.AddConfigPath(".")
	projectViper.AddConfigPath("../")
	projectViper.AddConfigPath("../../")
	projectViper.AddConfigPath("../../../")
	projectViper.SetConfigName(defaultConfigFileName)
	if err := projectViper.ReadInConfig(); nil != err {
		panic(err)
	}
	//股票信息配置
	var P *model.StockConfigs
	if err := projectViper.UnmarshalKey("stock", &P); nil != err {
		panic(err)
	}
	for _, stockConfig := range P.Configs {
		stockConfigMap[stockConfig.StockExchangeCode] = stockConfig
	}
	dataSaveFilePath = P.DataSavePath
	skipIfNoData = P.SkipIfNoData
}

func GetStockConfig(s int) model.StockConfig {
	return stockConfigMap[s]
}

func GetDataSaveFilePath() string {
	return dataSaveFilePath
}

func SkipNoData() bool {
	return skipIfNoData
}
