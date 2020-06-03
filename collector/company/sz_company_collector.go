package company

import (
    "../../config"
    "../../enums"
    "../../model"
    "../../s-logger"
    "encoding/json"
    "io/ioutil"
    "math/rand"
    "net/http"
    "regexp"
    "strconv"
    "strings"
)

var simpleNameReg = regexp.MustCompile("(.*)<u>(.*)</u>(.*)")
var logger = s_logger.New()

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

func (s SZCompanyCollector) String() string {
    return "SZCompanyCollector"
}

func (sz SZCompanyCollector) GetStockExchange() int {
    return enums.SZ
}

func (sz SZCompanyCollector) FetchAll(conf config.StockConfig) []model.Company {
    logger.Infow("收集所有公司信息,开始.", "stockExchangeCode", sz.GetStockExchange(), "configInfo", conf)
    result := make([]model.Company, 0)
    allPlate := enums.GetAll()
    for _, plate := range allPlate {
        result = append(result, GetPlateData(conf, plate)...)
    }
    logger.Infow("收集所有公司信息,结束.", "stockExchangeCode", sz.GetStockExchange(), "configInfo", conf, "length", len(result))
    return result
}

//获取每个板块的数据
func GetPlateData(conf config.StockConfig, plate enums.PlateEnum) []model.Company {
    logger.Infow("收集所有公司信息,收集指定板块信息,开始.", "stockExchangeCode", conf.StockExchangeCode, "plate", plate)
    result := make([]model.Company, 0)
    page := 1
    for pageData := readPageData(conf, page, plate); pageData != nil; {
        result = append(result, pageData...)
        page = page + 1
        pageData = readPageData(conf, page, plate)
    }
    logger.Infow("收集所有公司信息,收集指定板块信息,结束.", "stockExchangeCode", conf.StockExchangeCode, "plate", plate)
    return result
}

//读取每页的数据
func readPageData(conf config.StockConfig, page int, plate enums.PlateEnum) []model.Company {
    requestUrl := conf.CompanyInfoUrl + "&TABKEY=" + plate.Tab + "&random=" + strconv.Itoa(rand.Int()) + "&PAGENO=" + strconv.Itoa(page)
    logger.Infow("获取公司列表.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl)
    response, err := http.Get(requestUrl)
    if nil != err {
        logger.Errorw("获取公司列表数据异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
        return nil
    }
    responseDataByte, err := ioutil.ReadAll(response.Body)
    if nil != err {
        logger.Errorw("读取公司列表数据异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
        return nil
    }
    responseDataPoint := &[]Response{}
    err = json.Unmarshal(responseDataByte, &responseDataPoint)
    if err != nil {
        logger.Errorw("读取公司列表,解析数据出现异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
        return nil
    }

    if r := (*responseDataPoint)[plate.Index-1]; len(r.Data) > 0 {
        return responseToCompanyMapper(r, plate)
    }
    return nil
}

func responseToCompanyMapper(response Response, plate enums.PlateEnum) []model.Company {
    result := make([]model.Company, 0)
    for _, d := range response.Data {
        split := strings.Split(d.IndustryCodeAndName, " ")
        company := model.Company{
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
    logger.Infow("数据转换", "length", len(response.Data), "resultLength", len(result))
    return result
}

func getSimpleName(htmlStr string) string {
    return string(simpleNameReg.FindAllSubmatch([]byte(htmlStr), 1)[0][2])
}
