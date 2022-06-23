package data

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/xiaotian/stock/config"
	"github.com/xiaotian/stock/enums"
	"github.com/xiaotian/stock/model"
	"testing"
)

func TestSHFetchOne(t *testing.T) {

	conf := config.GetStockConfig(enums.SH)
	company := &[]model.Company{}

	var companyJson = "[{\"stock_exchange\": 130,\"code\": \"600000\",\"plate\": \"1000\",\"short_name\": \"浦发银行\",\"full_name\": \"-\",\"industry_code\": \"-\",\"industry_name\": \"-\"}]"
	err := json.Unmarshal([]byte(companyJson), company)
	assert.Nil(t, err, err)

	collector := SHDataCollector{}
	Data := collector.FetchAll(*company, conf)

	logger.Infow("Data", Data)
	assert.Equal(t, 1, len(Data))
}
