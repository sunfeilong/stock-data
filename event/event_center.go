package event

import (
    "../listener"
    "../s-logger"
    "./observer"
    "fmt"
    "sync"
    "time"
)

var logger = s_logger.New()

var EC *ECenter = &ECenter{observer: observer.New(), eventChain: make(chan Event, 100), finishedFlag: make(chan bool, 1)}

type ECenter struct {
    sync.Mutex                //锁
    waitGroup  sync.WaitGroup //等待任务完成

    observer *observer.Observer // 观察者

    eventChain   chan Event //事件管道
    finishedFlag chan bool  //完成标记

    totalEventCount      int //接收的事件总数量
    hasHandelEventCount  int //已经处理事件数量
    waitHandelEventCount int //等待处理事件数量
}

func HandelEvent() {
    logger.Infow("事件处理中心,启动处理程序.")
    go func() {
        for {
            event := <-EC.eventChain
            time.Sleep(time.Millisecond * 100)
            EC.Lock()
            EC.waitHandelEventCount = EC.waitHandelEventCount - 1
            EC.hasHandelEventCount = EC.hasHandelEventCount + 1
            EC.Unlock()
            logger.Infow("事件处理中心,获取到事件,开始处理.", "event", event, "ec", EC)
            EC.waitGroup.Done()
        }
    }()
}

func AddEvent(event Event) {
    EC.waitGroup.Add(1)
    EC.Lock()
    EC.eventChain <- event
    EC.totalEventCount = EC.totalEventCount + 1
    EC.waitHandelEventCount = EC.waitHandelEventCount + 1
    EC.Unlock()
    logger.Infow("事件处理中心,添加事件.", "event", event, "ec", EC)
}

func RegisterListener(listener listener.Listener) {
    if err := EC.observer.RegisterListener(listener); nil != err {
        logger.Errorw("注册监听者异常.", "err", err, "listener", listener)
    }
}

func RemoveListener(listener listener.Listener) {
    if err := EC.observer.RemoveListener(listener); nil != err {
        logger.Errorw("删除听者异常", "err", err, "listener", listener)
    }
}

func WaitEventHandleFinished() {
    logger.Infow("事件处理中心,等待事件处理完成,等待中.", "ec", EC)
    EC.waitGroup.Wait()
    logger.Infow("事件处理中心,等待事件处理完成.已完成", "ec", EC)
}

func (e ECenter) String() string {
    return fmt.Sprintf("totalEventCount:%0d, hasHandelEventCount:%0d,waitHandelEventCount:%0d",
        e.totalEventCount, e.hasHandelEventCount, e.waitHandelEventCount)
}
