package event

import "fmt"

type Code int

const (
    CollectCompanyInfoFinished Code = 1 + iota
)

//事件定义
type Event interface {
    GetType() string          //获取时间类型
    GetDescription() string   //获取时间描述
    GetMeteData() interface{} //获取事件元数据
    String() string
}

//事件
type EModel struct {
    eType        string      //事件类型
    eDescription string      //事件描述
    eMetaData    interface{} //事件元数据
}

func NewModel(eventType string, eventDescription string, metaData interface{}) EModel {
    return EModel{eType: eventType, eDescription: eventDescription, eMetaData: metaData}
}

func (m EModel) GetType() string {
    return m.eType
}

func (m EModel) GetDescription() string {
    return m.eDescription
}

func (m EModel) GetMeteData() interface{} {
    return m.eMetaData
}

func (m EModel) String() string {
    return fmt.Sprintf("type:%s, description:%s, metadata:%s", m.GetType(), m.GetDescription(), m.GetMeteData())
}
