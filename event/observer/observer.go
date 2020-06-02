package observer

import (
    "../../listener"
    "errors"
    "fmt"
)
import "../../s-logger"

var logger = s_logger.New()

func New() *Observer {
    return &Observer{observers: make(map[string][]listener.Listener)}
}

//观察者
type Observer struct {
    observers map[string][]listener.Listener
}

//增加监听者
func (o *Observer) RegisterListener(listener listener.Listener) error {
    logger.Infow("注册监听者,开始", "observers", o, "listener", listener)
    index := getIndex(listener, o)
    if index > 0 {
        logger.Infow("注册监听者,监听者已经存在", "observers", o, "listener ", listener)
        return errors.New("注册监听者,监听者已经存在")
    }
    o.observers[listener.GetEventType()] = append(o.observers[listener.GetEventType()], listener)
    logger.Infow("注册监听者,结束", "observers", o, "listener: ", listener)
    return nil
}

//删除监听者
func (o *Observer) RemoveListener(listener listener.Listener) error {
    logger.Infow("删除监听者,开始执行.", "observers", o, "listener", listener)
    listeners := o.observers[listener.GetEventType()]
    index := getIndex(listener, o)
    if index < 0 {
        logger.Infow("删除监听者,没有找打对应的监听者.", "observers", o, "listener", listener)
        return errors.New("删除监听者,没有找打对应的监听者")
    }
    copy(listeners[index:], listeners[index+1:])
    o.observers[listener.GetEventType()] = listeners[:len(listeners)-1]
    logger.Infow("删除监听者,删除成功.", "observers", o, "listener", listener)
    return nil
}

//获取监听者做标
func getIndex(listener listener.Listener, observer *Observer) int {
    listeners := observer.observers[listener.GetEventType()]
    for i, l := range listeners {
        if l.GetId() == listener.GetEventType() && l.GetEventType() == listener.GetEventType() {
            return i
        }
    }
    return -1
}

func (o Observer) String() string {
    return fmt.Sprint("observers: ", o.observers)
}
