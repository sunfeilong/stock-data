package enums

type PlateEnum struct {
    StockExchange StockExchange
    Code          string
    Name          string
    Tab           string
}

var MainPlate = PlateEnum{StockExchange: SZ, Code: "100", Name: "主板", Tab: "tab2"}
var MiddleOrLittlePlate = PlateEnum{StockExchange: SZ, Code: "200", Name: "中小企业版", Tab: "tab3"}
var PioneerPlatePlate = PlateEnum{StockExchange: SZ, Code: "300", Name: "创业版", Tab: "tab4"}

var result = []PlateEnum{MainPlate, MiddleOrLittlePlate, PioneerPlatePlate}

func getAll(enum PlateEnum) []PlateEnum {
    return result
}

func (p PlateEnum) codeToName(code string) string {
    for _, r := range result {
        if r.Code == code {
            return r.Name
        }
    }
    return ""

}

func (p PlateEnum) nameToCode(name string) string {
    for _, r := range result {
        if r.Name == name {
            return r.Code
        }
    }
    return ""
}
