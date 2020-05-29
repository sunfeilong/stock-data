package persistent

import (
    "../model"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestName(t *testing.T) {
    d := []model.Company{}
    d = append(d, model.Company{Code: "1"})
    d = append(d, model.Company{Code: "2"})

    cfp := CompanyFilePersistent{}

    err := cfp.save(d)

    assert.Nil(t, err, "")

    read := cfp.read()
    assert.NotEmpty(t, read, "")

}
