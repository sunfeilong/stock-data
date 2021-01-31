package company

import (
    "github.com/stretchr/testify/assert"
    "github.com/xiaotian/stock/pkg/config"
    "github.com/xiaotian/stock/pkg/enums"
    "testing"
)

func TestHKCompanyDataFetch(t *testing.T) {
    shCollector := HKCompanyCollector{}

    getConfig := config.GetStockConfig(enums.HK)

    companies := shCollector.FetchAll(getConfig)

    assert.NotEmpty(t, companies, "")
}
