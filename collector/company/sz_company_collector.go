package main

import (
    "../../config"
    "../../data"
    "../../enums"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "strconv"
)

//深圳交易所上市公司信息收集器
type SZCompanyCollector struct {
}

type Response struct {
    Data  []ResponseData `json:"data"`
    Error string         `json:"error"`
}

type ResponseData struct {
    CompanySimpleName   string `json:"gsjc"`
    CompanyFullName     string `json:"gsqc"`
    CompanyWebSite      string `json:"http"`
    IndustryCodeAndName string `json:"sshymc"`
    StockCode           string `json:"zqdm"`
}

func (sz SZCompanyCollector) getStockExchange() enums.StockExchange {
    return enums.SZ
}

func (sz SZCompanyCollector) fetchAll(conf config.Config) []data.Company {
    readPageData(conf, 1, enums.MainPlate)
    return nil
}

func readPageData(conf config.Config, page int, plate enums.PlateEnum) []data.Company {
    requestUrl := conf.CompanyInfoUrl + "&TABKEY=" + plate.Tab + "&random=" + strconv.Itoa(rand.Int()) + "&PAGENO=" + strconv.Itoa(page)
    log.Println("获取公司列表.交易所:", plate.StockExchange, ",板块", plate.Tab, "完整URL:%v", requestUrl)
    response, err := http.Get(requestUrl)
    if nil != err {
        log.Fatal("获取公司列表数据异常,异常信息", err)
    }
    log.Println("获取公司列表，响应信息: ", response)
    responseDataByte, err := ioutil.ReadAll(response.Body)
    if nil != err {
        log.Fatal("读取公司列表数据异常,异常信息: ", err)
    }
    log.Println("获取公司列表，响应数据:", string(responseDataByte))

    responseDataPoint := &[]Response{}
    err = json.Unmarshal(responseDataByte, &responseDataPoint)
    log.Println("获取公司列表，响应数据", responseDataPoint)
    return []data.Company{}
}

func main() {
    configInfo := config.YamlConfig{}
    szCollector := SZCompanyCollector{}
    getConfig, err := configInfo.GetConfig(enums.SZ)
    if nil != err {

    }
    companies := szCollector.fetchAll(getConfig)
    fmt.Println(companies)
}
