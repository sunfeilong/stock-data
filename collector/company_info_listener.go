package collector

import (
    "../model"
    "fmt"
)

func NewListener(id string, eventType string) CompanyInfoCollectorFinishedListener {
    c := make(chan *[]model.Company, 1)
    return CompanyInfoCollectorFinishedListener{id: id, eventType: eventType, data: c}
}

type CompanyInfoCollectorFinishedListener struct {
    id        string
    eventType string
    data      chan *[]model.Company
}

func (l CompanyInfoCollectorFinishedListener) GetId() string {
    return l.id
}

func (l CompanyInfoCollectorFinishedListener) GetEventType() string {
    return l.eventType
}

func (l CompanyInfoCollectorFinishedListener) Notify() {
    logger.Infow("处理事件")
}

func (l CompanyInfoCollectorFinishedListener) String() string {
    return fmt.Sprint("id: ", l.id, " eventType: ", l.eventType)
}

func (l *CompanyInfoCollectorFinishedListener) AddData(data *[]model.Company) {
    l.data <- data
}
