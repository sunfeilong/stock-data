package company

import (
    "../../config"
    "../../enums"
    "github.com/stretchr/testify/assert"
    "log"
    "testing"
)

func TestName(t *testing.T) {
    szCollector := SZCompanyCollector{}

    getConfig := config.GetStockConfig(enums.SZ)

    companies := szCollector.fetchAll(getConfig)
    log.Println("数据: ", companies)
    log.Println("数据长度: ", len(companies))
    assert.NotEmpty(t, companies, "")
}
