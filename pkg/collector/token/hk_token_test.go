package token

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestGetHKToken(t *testing.T) {
    token := GetHKToken("https://www.hkex.com.hk/?sc_lang=zh-HK")
    assert.NotEmpty(t, token, "GET TOKEN FAILED")
}
