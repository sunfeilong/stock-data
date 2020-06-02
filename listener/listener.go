package listener

//事件监听器
type Listener interface {
    GetId() string
    GetEventType() string
    String() string
    Handle()
}
