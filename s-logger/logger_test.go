package s_logger

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestName(t *testing.T) {
    logger := New()

    logger.Info("Info")
    logger.Error("Error")
    defer logger.Sync()
    assert.NotEmpty(t, logger, "日志信息不能为空")
}
