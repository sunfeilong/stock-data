package collector

import (
    "fmt"
)

func NewListener(id string, eventType string) CompanyInfoCollectorFinishedListener {
    return CompanyInfoCollectorFinishedListener{id: id, eventType: eventType}
}

type CompanyInfoCollectorFinishedListener struct {
    id        string
    eventType string
}

func (l CompanyInfoCollectorFinishedListener) GetId() string {
    return l.id
}

func (l CompanyInfoCollectorFinishedListener) GetEventType() string {
    return l.eventType
}

func (l CompanyInfoCollectorFinishedListener) Handle() {
    logger.Infow("处理事件")
}

func (l CompanyInfoCollectorFinishedListener) String() string {
    return fmt.Sprint("id: ", l.id, " eventType: ", l.eventType)
}
