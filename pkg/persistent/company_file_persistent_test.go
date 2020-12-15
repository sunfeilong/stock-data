package persistent

import (
    "github.com/stretchr/testify/assert"
    "github.com/xiaotian/stock/pkg/model"
    "testing"
)

func TestName(t *testing.T) {
    d := make([]model.Company, 0)
    d = append(d, model.Company{Code: "1"})
    d = append(d, model.Company{Code: "2"})

    cfp := CompanyFilePreserver{}

    err := cfp.Save(d)

    assert.Nil(t, err, "")

    read, err := cfp.Read()
    assert.Empty(t, err)
    assert.NotEmpty(t, read, "")
    assert.Equal(t, "1", read[0].Code)
    assert.Equal(t, "2", read[1].Code)
}
