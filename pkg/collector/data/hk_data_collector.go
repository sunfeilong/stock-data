package data

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/xiaotian/stock/pkg/collector/token"
    "github.com/xiaotian/stock/pkg/enums"
    "github.com/xiaotian/stock/pkg/model"
    "io/ioutil"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"
)

type HKDataCollector struct {
}

type HKResponse struct {
    Data HKResponseData `json:"sample"`
}

type HKResponseData struct {
    ResponseCode string          `json:"responsecode"`
    ResponseMsg  string          `json:"responsemsg"`
    DataList     [][]interface{} `json:"datalist"`
}

func (s HKDataCollector) GetStockExchange() int {
    return enums.HK
}

func (s HKDataCollector) FetchAll(company []model.Company, conf model.StockConfig) []model.Data {
    logger.Infow("获取港交所上市公司股票价格数据,", "company count", len(company), "conf", conf)
    result := make([]model.Data, 0)
    hkToken := token.GetHKToken(conf.TokenUrl)
    for _, c := range company {
        if s.GetStockExchange() != c.StockExchange {
            logger.Infow("获取港交所上市公司股票价格数据,收集器不能处理对应公司数据,跳过",
                "company", c, "conf", conf,
                "collectorStockExchange", c.StockExchange, "companyStockExchange", c.StockExchange)
            continue
        }
        logger.Infow("获取港交所上市公司股票价格数据", "company", c, "conf", conf)
        data, err := HKGetData(c, conf, hkToken)
        if err != nil {
            logger.Errorw("获取港交所上市公司股票价格数据,获取数据异常", "company", c, "conf", conf, "err", err)
            continue
        }
        result = append(result, data)
    }
    return result
}

func HKGetData(company model.Company, config model.StockConfig, hkToken string) (model.Data, error) {
    time.Sleep(time.Millisecond * 500)
    data := &model.Data{}
    data.StockExchange = company.StockExchange
    data.Code = company.Code
    data.Plate = company.Plate
    url := strings.Replace(config.RealTimeInfoUrl, "{code}", fmt.Sprintf("%04s", company.Code), -1)
    url = strings.Replace(url, "{token}", hkToken, -1)
    url = url + "&qid=" + strconv.Itoa(rand.Int()) + "&_=" + strconv.Itoa(rand.Int())
    logger.Infow("获取港交所上市公司股票数据,开始", "code", company.Code, "company", company.ShortName, "url", url)
    client := &http.Client{}
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        logger.Errorw("获取港交所上市公司股票数据,构造请求出现异常", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
        return *data, errors.New("构造请求出现异常")
    }
    request.Header.Add("Cookie", COOKIE)
    request.Header.Add("User-Agent", USERAGENT)
    request.Header.Add("Referer", Referer)
    response, err := client.Do(request)
    if err != nil {
        logger.Errorw("获取港交所上市公司股票数据,请求数据出现异常", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
        return *data, errors.New("请求数据出现异常")
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        logger.Errorw("获取港交所上市公司股票数据,读取响应数据出错", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
        return *data, errors.New("读取响应数据出错")
    }
    tempStr := string(responseData)
    tempStr = tempStr[strings.Index(tempStr, "(")+1 : strings.LastIndex(tempStr, ")")]
    responseDataPointer := &HKResponse{}
    if err = json.Unmarshal([]byte(tempStr), responseDataPointer); err != nil {
        logger.Errorw("获取港交所上市公司股票数据,解析数据出错", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
        return *data, errors.New("解析数据出错")
    }
    HKCopyData(data, responseDataPointer)
    logger.Infow("获取港交所上市公司股票数据,结束", "code", company.Code, "company", company.ShortName, "url", url, )
    return *data, nil
}

func HKCopyData(data *model.Data, response *HKResponse) {
    for _, p := range response.Data.DataList {
        if p[1] == nil{
            continue
        }
        unixTime := time.Unix(int64(p[0].(float64))/1000, 0).Format("2006-01-02 15:04")
        data.AddInnerData(unixTime, fmt.Sprintf("%.2f", p[4]))
    }
}
