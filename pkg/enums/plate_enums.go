package enums

import "github.com/xiaotian/stock/pkg/config"

type PlateEnum struct {
    StockExchange int
    Code          string
    Name          string
    Tab           string
    Index         int
}

var MainPlate = PlateEnum{StockExchange: SZ, Code: "100", Name: "主板", Tab: "tab2", Index: 2}
var MiddleOrLittlePlate = PlateEnum{StockExchange: SZ, Code: "200", Name: "中小企业版", Tab: "tab3", Index: 3}
var PioneerPlatePlate = PlateEnum{StockExchange: SZ, Code: "300", Name: "创业版", Tab: "tab4", Index: 4}
var APlate = PlateEnum{StockExchange: SH, Code: "1000", Name: "A股", Tab: "1", Index: 0}
var BPlate = PlateEnum{StockExchange: SH, Code: "1100", Name: "B股", Tab: "2", Index: 0}
var TechnologyPlate = PlateEnum{StockExchange: SH, Code: "1200", Name: "科创版", Tab: "8", Index: 0}

var result = []PlateEnum{MainPlate, MiddleOrLittlePlate, PioneerPlatePlate, APlate, BPlate, TechnologyPlate}

func GetAll() []PlateEnum {
    return result
}

func GetByStockExchange(sc config.StockConfig) []PlateEnum {
    temp := make([]PlateEnum, 0)
    for _, r := range result {
        if r.StockExchange == sc.StockExchangeCode {
            temp = append(temp, r)
        }
    }
    return temp
}

func (p PlateEnum) CodeToName(code string) string {
    for _, r := range result {
        if r.Code == code {
            return r.Name
        }
    }
    return ""

}

func (p PlateEnum) NameToCode(name string) string {
    for _, r := range result {
        if r.Name == name {
            return r.Code
        }
    }
    return ""
}
