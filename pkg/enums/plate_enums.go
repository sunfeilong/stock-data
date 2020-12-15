package enums

import "github.com/xiaotian/stock/pkg/config"

type PlateEnum struct {
    StockExchange int
    Code          string
    Name          string
    Tab           string
    Index         int
}

var SZMainPlate = PlateEnum{StockExchange: SZ, Code: "100", Name: "主板", Tab: "tab2", Index: 2}
var SZMiddleOrLittlePlate = PlateEnum{StockExchange: SZ, Code: "200", Name: "中小企业版", Tab: "tab3", Index: 3}
var SZPioneerPlatePlate = PlateEnum{StockExchange: SZ, Code: "300", Name: "创业版", Tab: "tab4", Index: 4}
var SHAPlate = PlateEnum{StockExchange: SH, Code: "1000", Name: "A股", Tab: "1", Index: 0}
var SHBPlate = PlateEnum{StockExchange: SH, Code: "1100", Name: "B股", Tab: "2", Index: 0}
var SHTechnologyPlate = PlateEnum{StockExchange: SH, Code: "1200", Name: "科创版", Tab: "8", Index: 0}
var HKMainPlate = PlateEnum{StockExchange: HK, Code: "2000", Name: "主板", Tab: "MAIN", Index: 0}
var HKGemPlate = PlateEnum{StockExchange: HK, Code: "2100", Name: "GEM", Tab: "GEM", Index: 0}

var result = []PlateEnum{SZMainPlate, SZMiddleOrLittlePlate, SZPioneerPlatePlate, SHAPlate, SHBPlate, SHTechnologyPlate, HKMainPlate, HKGemPlate}

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
