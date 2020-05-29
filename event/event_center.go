package event

import (
    "../s-logger"
    "fmt"
    "sync"
    "time"
)

var logger = s_logger.New()

type ECenter struct {
    sync.Mutex                      //锁
    eventChain           chan Event //事件管道
    totalEventCount      int        //接收的事件总数量
    hasHandelEventCount  int        //已经处理事件数量
    waitHandelEventCount int        //等待处理事件数量
    finishedFlag         chan bool  //完成标记
}

func NewECenter(queueSize int) *ECenter {
    return &ECenter{eventChain: make(chan Event, queueSize), finishedFlag: make(chan bool, 1)}
}

func (e *ECenter) HandelEvent() {
    logger.Infow("事件处理中心,启动处理程序.")
    go func() {
        for {
            event := <-e.eventChain
            time.Sleep(time.Millisecond * 100)
            e.Lock()
            e.waitHandelEventCount = e.waitHandelEventCount - 1
            e.hasHandelEventCount = e.hasHandelEventCount + 1
            e.Unlock()
            logger.Infow("事件处理中心,获取到事件,开始处理.", "event", event, "ec", e)
            if e.waitHandelEventCount == 0 {
                e.finishedFlag <- true
            }
        }
    }()
}

func (e *ECenter) AddEvent(event Event) {
    e.Lock()
    e.eventChain <- event
    e.totalEventCount = e.totalEventCount + 1
    e.waitHandelEventCount = e.waitHandelEventCount + 1
    e.Unlock()
    logger.Infow("事件处理中心,添加事件.", "event", event, "ec", e)
}

func (e *ECenter) WaitFinished() {
    logger.Infow("事件处理中心,等待处理完成.", "state", e)
    <-e.finishedFlag
    logger.Infow("事件处理中心,处理已经完成.", "state", e)
}

func (e *ECenter) String() string {
    return fmt.Sprintf("totalEventCount:%d, hasHandelEventCount:%d,waitHandelEventCount:%d", e.totalEventCount, e.hasHandelEventCount, e.waitHandelEventCount)
}
