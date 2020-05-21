package company

import (
    "../../config"
    "../../data"
    "../../enums"
    "encoding/json"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "regexp"
    "strconv"
    "strings"
)

var simpleNameReg = regexp.MustCompile("(.*)<u>(.*)</u>(.*)")

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
    result := make([]data.Company, 0)
    allPlate := enums.GetAll()
    for _, plate := range allPlate {
        plateData := getPlateData(conf, plate)
        result = append(result, plateData...)
    }
    return result
}

//获取每个板块的数据
func getPlateData(conf config.Config, plate enums.PlateEnum) []data.Company {
    result := make([]data.Company, 0)
    page := 1
    for pageData := readPageData(conf, page, plate); pageData != nil; {
        result = append(result, pageData...)
        page = page + 1
        pageData = readPageData(conf, page, plate)
    }
    return result
}

//读取每页的数据
func readPageData(conf config.Config, page int, plate enums.PlateEnum) []data.Company {
    requestUrl := conf.CompanyInfoUrl + "&TABKEY=" + plate.Tab + "&random=" + strconv.Itoa(rand.Int()) + "&PAGENO=" + strconv.Itoa(page)
    log.Println("获取公司列表.交易所:", plate.StockExchange, ",板块", plate.Tab, "完整URL:%v", requestUrl)
    response, err := http.Get(requestUrl)
    if nil != err {
        log.Fatal("获取公司列表数据异常,异常信息", err)
    }
    //log.Println("获取公司列表，响应信息: ", response)
    responseDataByte, err := ioutil.ReadAll(response.Body)
    if nil != err {
        log.Fatal("读取公司列表数据异常,异常信息: ", err)
    }
    //log.Println("获取公司列表，响应数据:", string(responseDataByte))
    responseDataPoint := &[]Response{}
    err = json.Unmarshal(responseDataByte, &responseDataPoint)

    var r Response
    for index, d := range *responseDataPoint {
        if index == plate.Index-1 {
            r = d
        }
    }
    if len(r.Data) == 0 {
        return nil
    }
    //log.Println("获取公司列表，响应数据", responseDataPoint)
    return responseToCompanyMapper(r, plate)
}

func responseToCompanyMapper(response Response, plate enums.PlateEnum) []data.Company {
    result := make([]data.Company, 0)
    for _, d := range response.Data {
        split := strings.Split(d.IndustryCodeAndName, " ")
        company := data.Company{
            StockExchange: plate.StockExchange,
            Code:          d.StockCode,
            Plate:         plate.Code,
            ShortName:     getSimpleName(d.CompanySimpleName),
            FullName:      d.CompanyFullName,
            IndustryCode:  split[0],
            IndustryName:  split[1],
        }
        result = append(result, company)
    }
    log.Println("数据量:", len(response.Data), "转换结果数量: ", len(result))
    return result
}

func getSimpleName(htmlStr string) string {
    return string(simpleNameReg.FindAllSubmatch([]byte(htmlStr), 1)[0][2])
}
