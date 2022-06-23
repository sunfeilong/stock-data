package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xiaotian/stock/enums"
	"github.com/xiaotian/stock/model"
	"github.com/xiaotian/stock/tool"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	COOKIE    string = "yfx_c_g_u_id_10000042=_ck18012900250116338392357618947; VISITED_MENU=%5B%228528%22%5D; yfx_f_l_v_t_10000042=f_t_1517156701630__r_t_1517314287296__v_t_1517320502571__r_c_2"
	USERAGENT string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36"
	Referer   string = "http://www.sse.com.cn/assortment/stock/list/share/"
)

type SHDataCollector struct {
}

type SHResponse struct {
	Data [][]interface{} `json:"line"`
}

func (s SHDataCollector) GetStockExchange() int {
	return enums.SH
}

func (s SHDataCollector) FetchAll(company []model.Company, conf model.StockConfig) []model.Data {
	logger.Infow("获取上交所上市公司股票价格数据,", "company count", len(company), "conf", conf)
	result := make([]model.Data, 0)
	for _, c := range company {
		if s.GetStockExchange() != c.StockExchange {
			logger.Infow("获取上交所上市公司股票价格数据,收集器不能处理对应公司数据,跳过",
				"company", c, "conf", conf,
				"collectorStockExchange", c.StockExchange, "companyStockExchange", c.StockExchange)
			continue
		}
		logger.Infow("获取上交所上市公司股票价格数据", "company", c, "conf", conf)
		data, err := SHGetData(c, conf)
		if err != nil {
			logger.Errorw("获取上交所上市公司股票价格数据,获取数据异常", "company", c, "conf", conf, "err", err)
			continue
		}
		result = append(result, data)
	}
	return result
}

func SHGetData(company model.Company, config model.StockConfig) (model.Data, error) {
	time.Sleep(time.Millisecond * 100)
	data := &model.Data{}
	data.StockExchange = company.StockExchange
	data.Code = company.Code
	data.Plate = company.Plate

	url := strings.Replace(config.RealTimeInfoUrl, "{code}", company.Code, -1)
	logger.Infow("获取上交所上市公司股票数据,开始", "code", company.Code, "company", company.ShortName, "url", url)

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Errorw("获取上交所上市公司股票数据,构造请求出现异常", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
		return *data, errors.New("构造请求出现异常")
	}
	request.Header.Add("Cookie", COOKIE)
	request.Header.Add("User-Agent", USERAGENT)
	request.Header.Add("Referer", Referer)
	response, err := client.Do(request)
	if err != nil {
		logger.Errorw("获取上交所上市公司股票数据,请求数据出现异常", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
		return *data, errors.New("请求数据出现异常")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Errorw("获取上交所上市公司股票数据,读取响应数据出错", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
		return *data, errors.New("读取响应数据出错")
	}
	responseDataPointer := &SHResponse{}
	if err = json.Unmarshal(responseData, responseDataPointer); err != nil {
		logger.Errorw("获取上交所上市公司股票数据,解析数据出错", "code", company.Code, "company", company.ShortName, "url", url, "error", err)
		return *data, errors.New("解析数据出错")
	}
	SHCopyData(data, responseDataPointer)
	logger.Infow("获取上交所上市公司股票数据,结束", "code", company.Code, "company", company.ShortName, "url", url)
	return *data, nil
}

func SHCopyData(data *model.Data, response *SHResponse) {
	for _, p := range response.Data {
		t := fmt.Sprintf("%.0f", p[0])
		if len(t) == 5 {
			t = "0" + t
		}
		data.AddInnerData(tool.NowDate()+" "+t[:2]+":"+t[2:4], fmt.Sprintf("%.0f", p[1]))
	}
}
