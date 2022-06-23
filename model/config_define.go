package model

import (
    "fmt"
)

//配置信息定义
type StockConfig struct {
    StockExchangeCode int
    StockExchange     string
    CompanyInfoUrl    string
    StockInfoUrl      string
    RealTimeInfoUrl   string
    TokenUrl          string
}

type StockConfigs struct {
    Configs      []StockConfig
    DataSavePath string
    SkipIfNoData bool
}

func (s StockConfig) String() string {
    return fmt.Sprintf("StockExchangeCode: %d, StockExchange: %s, CompanyInfoUrl: %s, StockInfoUrl: %s, RealTimeInfoUrl: %s",
        s.StockExchangeCode, s.StockExchange, s.CompanyInfoUrl, s.StockInfoUrl, s.RealTimeInfoUrl)
}
