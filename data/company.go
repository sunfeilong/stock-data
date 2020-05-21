package data

import "../enums"

//股票公司
type Company struct {
    StockExchange enums.StockExchange
    Code          string
    Plate         string
    ShortName     string
    FullName      string
    IndustryCode  string
    IndustryName  string
}
