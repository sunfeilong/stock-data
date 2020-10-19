package data

import (
    "encoding/json"
    "errors"
    "github.com/xiaotian/stock/pkg/config"
    "github.com/xiaotian/stock/pkg/enums"
    "github.com/xiaotian/stock/pkg/model"
    "github.com/xiaotian/stock/pkg/s-logger"
    "github.com/xiaotian/stock/pkg/tool"
    "io/ioutil"
    "net/http"
)

var logger = s_logger.New()

type SZDataCollector struct {
}

type Response struct {
    Code  string       `json:"code"`
    Data  ResponseData `json:"data"`
    Error string       `json:"error"`
}

type ResponseData struct {
    PicUpData [][]interface{} `json:"picupdata"`
}

func (s SZDataCollector) GetStockExchange() int {
    return enums.SZ
}

func (s SZDataCollector) FetchAll(company []model.Company, conf config.StockConfig) []model.Data {
    logger.Infow("获取深交所上市公司股票价格数据,", "company count", len(company), "conf", conf)
    result := make([]model.Data, 0)
    for _, c := range company {
        if s.GetStockExchange() != c.StockExchange {
            logger.Infow("获取深交所上市公司股票价格数据,收集器不能处理对应公司数据,跳过",
                "company", c, "conf", conf,
                "collectorStockExchange", c.StockExchange, "companyStockExchange", c.StockExchange)
            continue
        }
        logger.Infow("获取深交所上市公司股票价格数据", "company", c, "conf", conf)
        data, err := getData(c, conf)
        if err != nil {
            logger.Errorw("获取深交所上市公司股票价格数据,获取数据异常", "company", c, "conf", conf, "err", err)
            continue
        }
        result = append(result, data)
    }
    return result
}

func getData(company model.Company, config config.StockConfig) (model.Data, error) {
    data := &model.Data{}
    data.StockExchange = company.StockExchange
    data.Code = company.Code
    data.Plate = company.Plate

    url := config.RealTimeInfoUrl + "&code=" + company.Code
    logger.Infow("获取深交所上市公司股票数据,开始", "code", company.Code, "company", company.ShortName, "url", url)
    response, err := http.Get(url)
    if err != nil {
        logger.Errorw("获取深交所上市公司股票数据,请求数据出现异常", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
        return *data, errors.New("请求数据出现异常")
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        logger.Errorw("获取深交所上市公司股票数据,读取响应数据出错", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
        return *data, errors.New("读取响应数据出错")
    }
    responseDataPointer := &Response{}
    if err = json.Unmarshal(responseData, responseDataPointer); err != nil {
        logger.Errorw("获取深交所上市公司股票数据,解析数据出错", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
        return *data, errors.New("解析数据出错")
    }
    copyData(data, responseDataPointer)
    logger.Infow("获取深交所上市公司股票数据,结束", "code", company.Code, "company", company.ShortName, "url", url, )
    return *data, nil
}

func copyData(data *model.Data, response *Response) {
    for _, p := range response.Data.PicUpData {
        data.AddInnerData(tool.NowDate()+" "+p[0].(string), p[1].(string))
    }
}
