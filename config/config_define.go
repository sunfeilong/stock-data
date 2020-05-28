package config

//配置信息定义
type StockConfig struct {
    StockExchangeCode int
    StockExchange     string
    CompanyInfoUrl    string
    StockInfoUrl      string
    RealTimeInfoUrl   string
}

type StockConfigs struct {
    Configs []StockConfig
}

func (s StockConfig) String() string {
    return ""
}
