package persistent

import "github.com/xiaotian/stock/model"

//持久接口
type Preserver interface {
	Save(data []model.Company) error
	Read() ([]model.Company, error)
}
