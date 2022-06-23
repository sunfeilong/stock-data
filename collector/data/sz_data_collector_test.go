package data

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/xiaotian/stock/config"
	"github.com/xiaotian/stock/enums"
	"github.com/xiaotian/stock/model"
	"testing"
)

func TestFetchOne(t *testing.T) {

	conf := config.GetStockConfig(enums.SZ)
	company := &[]model.Company{}

	var companyJson = "[{\"stock_exchange\": 100,\"code\": \"000001\",\"plate\": \"100\",\"short_name\": \"平安银行\",\"full_name\": \"平安银行股份有限公司\",\"industry_code\": \"J\",\"industry_name\": \"金融业\"}]"
	err := json.Unmarshal([]byte(companyJson), company)
	assert.Nil(t, err, err)

	collector := SZDataCollector{}
	Data := collector.FetchAll(*company, conf)

	logger.Infow("Data", Data)
	assert.Equal(t, 1, len(Data))
}
