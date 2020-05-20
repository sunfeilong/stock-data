package data

import "../enums"

//股票公司
type Company struct {
    stockExchange enums.StockExchange
    code          string
    shortName     string
    fullName      string
    industryCode  string
    industryName  string
}
