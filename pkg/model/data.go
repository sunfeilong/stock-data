package model

type Data struct {
    StockExchange int         `json:"stock_exchange"`
    Code          string      `json:"code"`
    Plate         string      `json:"plate"`
    Data          []InnerData `json:"sample"`
}

func (d *Data) AddInnerData(date string, price string) {
    d.Data = append(d.Data, InnerData{Date: date, Price: price})
}

type InnerData struct {
    Date  string `json:"date"`
    Price string `json:"price"`
}
