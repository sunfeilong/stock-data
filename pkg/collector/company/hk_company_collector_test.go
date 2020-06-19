package company

import (
    "github.com/stretchr/testify/assert"
    "github.com/xiaotian/stock/pkg/config"
    "github.com/xiaotian/stock/pkg/enums"
    "log"
    "testing"
)

func TestHKCompanyDataFetch(t *testing.T) {
    shCollector := HKCompanyCollector{}

    getConfig := config.GetStockConfig(enums.HK)

    companies := shCollector.FetchAll(getConfig)
    log.Println("数据: ", companies)
    log.Println("数据长度: ", len(companies))
    assert.NotEmpty(t, companies, "")
}
