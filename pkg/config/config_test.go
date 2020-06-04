package config

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "github.com/xiaotian/stock/pkg/enums"
    "testing"
)

func TestGetStockExchangeConfig(t *testing.T) {
    HKConfig := GetStockConfig(enums.HK)

    assert.NotEmpty(t, HKConfig, "配置信息为空，没有获取到配置信息")
    fmt.Println(HKConfig)
    assert.Equal(t, "HK", HKConfig.StockExchange, "配置信息不正确")
}
