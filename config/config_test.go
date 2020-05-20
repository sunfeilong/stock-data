package config

import "testing"
import "github.com/stretchr/testify/assert"
import "../enums"

func TestGetStockExchangeConfig(t *testing.T) {
    c := YamlConfig{}
    HKConfig, eil := c.getConfig(enums.HK)

    assert.Empty(t, eil, "获取配置信息出错")
    assert.NotEmpty(t, HKConfig, "配置信息为空，没有获取到配置信息")
    assert.Equal(t, "HK", HKConfig.StockExchange, "配置信息不正确")
}
