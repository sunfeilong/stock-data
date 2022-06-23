package company

import (
	"encoding/json"
	"github.com/xiaotian/stock/collector/token"
	"github.com/xiaotian/stock/enums"
	"github.com/xiaotian/stock/model"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//香港交易所上市公司信息收集器
type HKCompanyCollector struct {
}

type HKResponse struct {
	Data  HKStockList `json:"data"`
	Error string      `json:"qid"`
}

type HKStockList struct {
	LastUpDate   string           `json:"lastupd"`
	ResponseCode string           `json:"responsecode"`
	ResponseMsg  string           `json:"responsemsg"`
	StockList    []HKResponseData `json:"stocklist"`
}

type HKResponseData struct {
	Suspend         bool   `json:"suspend"`
	CompanyFullName string `json:"nm"`
	StockCode       string `json:"sym"`
}

func (s HKCompanyCollector) String() string {
	return "HKCompanyCollector"
}

func (s HKCompanyCollector) GetStockExchange() int {
	return enums.HK
}

func (s HKCompanyCollector) FetchAll(conf model.StockConfig) []model.Company {
	logger.Infow("收集港交所所有公司信息,开始.", "stockExchangeCode", s.GetStockExchange(), "configInfo", conf)
	result := make([]model.Company, 0)
	allPlate := enums.GetByStockExchange(conf)
	hkToken := token.GetHKToken(conf.TokenUrl)
	for _, plate := range allPlate {
		result = append(result, HKGetPlateData(conf, plate, hkToken)...)
	}
	logger.Infow("收集港交所所有公司信息,结束.", "stockExchangeCode", s.GetStockExchange(), "configInfo", conf, "length", len(result))
	return result
}

//获取每个板块的数据
func HKGetPlateData(conf model.StockConfig, plate enums.PlateEnum, hkToken string) []model.Company {
	logger.Infow("收集所有公司信息,收集指定板块信息,开始.", "stockExchangeCode", conf.StockExchangeCode, "plate", plate)
	data := HKReadPageData(conf, plate, hkToken)
	logger.Infow("收集所有公司信息,收集指定板块信息,结束.", "stockExchangeCode", conf.StockExchangeCode, "plate", plate)
	return data
}

//读取每页的数据
func HKReadPageData(conf model.StockConfig, plate enums.PlateEnum, hkToken string) []model.Company {
	time.Sleep(time.Millisecond * 500)
	requestUrl := conf.CompanyInfoUrl + "&market=" + plate.Tab + "&qid=" + strconv.Itoa(rand.Int()) + "&_=" + strconv.Itoa(rand.Int())
	requestUrl = strings.ReplaceAll(requestUrl, "{token}", hkToken)
	logger.Infow("获取港交所公司列表.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl)
	response, err := http.Get(requestUrl)
	if nil != err {
		logger.Errorw("获取港交所公司列表数据异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
		return nil
	}
	responseDataByte, err := ioutil.ReadAll(response.Body)
	if nil != err {
		logger.Errorw("读取港交所公司列表数据异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
		return nil
	}
	responseDataPoint := HKResponse{}
	tempStr := string(responseDataByte)
	tempStr = tempStr[strings.Index(tempStr, "(")+1 : strings.LastIndex(tempStr, ")")]
	logger.Infow("读取港交所公司列表,解析数据", "str:", tempStr)
	err = json.Unmarshal([]byte(tempStr), &responseDataPoint)
	if err != nil {
		logger.Errorw("读取港交所公司列表,解析数据出现异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
		return nil
	}
	return HKResponseToCompanyMapper(responseDataPoint, plate)
}

func HKResponseToCompanyMapper(response HKResponse, plate enums.PlateEnum) []model.Company {
	result := make([]model.Company, 0)
	if response.Data.ResponseCode != "000" {
		return result
	}
	for _, d := range response.Data.StockList {
		if d.Suspend {
			logger.Infow("港交所公司列表数据转换,公司已停牌", "股票代码: ", d.StockCode, " 公司名字: ", d.CompanyFullName)
			continue
		}
		company := model.Company{
			StockExchange: plate.StockExchange,
			Code:          d.StockCode,
			Plate:         plate.Code,
			ShortName:     "-",
			FullName:      d.CompanyFullName,
			IndustryCode:  "-",
			IndustryName:  "-",
		}
		result = append(result, company)
	}
	logger.Infow("港交所公司列表数据转换", "length", len(response.Data.StockList), "resultLength", len(result))
	return result
}
