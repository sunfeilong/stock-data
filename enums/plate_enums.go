package enums

type PlateEnum struct {
    StockExchange StockExchange
    Code          string
    Name          string
    Tab           string
    Index         int
}

var MainPlate = PlateEnum{StockExchange: SZ, Code: "100", Name: "主板", Tab: "tab2", Index: 2}
var MiddleOrLittlePlate = PlateEnum{StockExchange: SZ, Code: "200", Name: "中小企业版", Tab: "tab3", Index: 3}
var PioneerPlatePlate = PlateEnum{StockExchange: SZ, Code: "300", Name: "创业版", Tab: "tab4", Index: 4}

var result = []PlateEnum{MainPlate, MiddleOrLittlePlate, PioneerPlatePlate}

func GetAll() []PlateEnum {
    return result
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
