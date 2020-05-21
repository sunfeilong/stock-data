package persistent

import (
    "../data"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestName(t *testing.T) {
    d := []data.Company{}
    d = append(d, data.Company{Code: "1"})
    d = append(d, data.Company{Code: "2"})

    cfp := CompanyFilePersistent{}

    err := cfp.save(d)

    assert.Nil(t, err, "")

    read := cfp.read()
    assert.NotEmpty(t, read, "")

}
