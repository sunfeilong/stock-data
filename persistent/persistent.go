package persistent

import "../model"

//持久接口
type Preserver interface {
    save(data []model.Company) error
    read() ([]model.Company, error)
}
