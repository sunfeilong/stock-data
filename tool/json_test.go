package tool

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestJSONToString(t *testing.T) {

    toString := JSONToString("123")
    assert.NotEmpty(t, toString)
    assert.Equal(t, "\"123\"", toString)

    toString = JSONToString(Data{Name: "123", age: "456"})

    assert.NotEmpty(t, toString)
    assert.Equal(t, "{\"Name\":\"123\"}", toString)
}

type Data struct {
    Name string
    age  string
}
