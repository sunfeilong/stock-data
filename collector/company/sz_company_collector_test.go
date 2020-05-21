package company

import (
    "../../config"
    "../../enums"
    "github.com/stretchr/testify/assert"
    "log"
    "testing"
)

func TestName(t *testing.T) {
    configInfo := config.YamlConfig{}
    szCollector := SZCompanyCollector{}

    getConfig, err := configInfo.GetConfig(enums.SZ)
    assert.Nil(t, err, "")

    companies := szCollector.fetchAll(getConfig)
    log.Println("数据: ", companies)
    log.Println("数据长度: ", len(companies))
    assert.NotEmpty(t, companies, "")
}
