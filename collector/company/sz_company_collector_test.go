package company

import (
	"github.com/stretchr/testify/assert"
	"github.com/xiaotian/stock/config"
	"github.com/xiaotian/stock/enums"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	szCollector := SZCompanyCollector{}

	getConfig := config.GetStockConfig(enums.SZ)

	companies := szCollector.FetchAll(getConfig)
	log.Println("数据: ", companies)
	log.Println("数据长度: ", len(companies))
	assert.NotEmpty(t, companies, "")
}
