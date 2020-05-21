package persistent

import "../data"

//持久接口
type Persistent interface {
    save(data []data.Company) error
    read() ([]data.Company, error)
}
