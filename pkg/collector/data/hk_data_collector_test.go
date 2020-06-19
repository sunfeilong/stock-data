package data

import (
    "encoding/json"
    "github.com/stretchr/testify/assert"
    "github.com/xiaotian/stock/pkg/config"
    "github.com/xiaotian/stock/pkg/enums"
    "github.com/xiaotian/stock/pkg/model"
    "testing"
)

func TestHKFetchAll(t *testing.T) {

    conf := config.GetStockConfig(enums.HK)
    company := &[]model.Company{}

    var companyJson = "[{\"stock_exchange\": 160,\"code\": \"700\",\"plate\": \"2000\",\"short_name\": \"腾讯\",\"full_name\": \"-\",\"industry_code\": \"-\",\"industry_name\": \"-\"}]"
    err := json.Unmarshal([]byte(companyJson), company)
    assert.Nil(t, err, err)

    collector := HKDataCollector{}
    Data := collector.FetchAll(*company, conf)

    logger.Infow("Data", Data)
    assert.Equal(t, 1, len(Data))
}
