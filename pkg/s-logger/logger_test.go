package s_logger

import (
    "github.com/stretchr/testify/assert"
    "sync"
    "testing"
)

func TestName(t *testing.T) {
    logger := New()

    logger.Info("Info")
    logger.Error("Error")
    defer logger.Sync()
    assert.NotEmpty(t, logger, "日志信息不能为空")
}

func TestMultiSingle(t *testing.T) {
    logger := New()
    times := 102400
    for i := 0; i < times; i++ {
        logger.Infow("测试打印日志", "name", "name")
    }
}

func TestMultiOpen(t *testing.T) {
    waitGroup := sync.WaitGroup{}
    waitGroup.Add(2)
    go func() {
        logger := New()
        times := 102400
        for i := 0; i < times; i++ {
            logger.Infow("1111111111", "name", "name")
        }
        waitGroup.Done()
    }()

    go func() {
        logger := New()
        times := 102400
        for i := 0; i < times; i++ {
            logger.Infow("2222222222", "name", "name")
        }
        waitGroup.Done()
    }()

    waitGroup.Wait()
}
