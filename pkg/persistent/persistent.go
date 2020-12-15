package persistent

import "github.com/xiaotian/stock/pkg/model"

//持久接口
type Preserver interface {
    Save(data []model.Company) error
    Read() ([]model.Company, error)
}
