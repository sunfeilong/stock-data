package company

import (
    "encoding/json"
    "github.com/xiaotian/stock/pkg/config"
    "github.com/xiaotian/stock/pkg/enums"
    "github.com/xiaotian/stock/pkg/model"
    "io/ioutil"
    "math/rand"
    "net/http"
    "strconv"
    "time"
)

//上海交易所上市公司信息收集器
type SHCompanyCollector struct {
}

const (
    COOKIE    string = "yfx_c_g_u_id_10000042=_ck18012900250116338392357618947; VISITED_MENU=%5B%228528%22%5D; yfx_f_l_v_t_10000042=f_t_1517156701630__r_t_1517314287296__v_t_1517320502571__r_c_2"
    USERAGENT string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36"
    Referer   string = "http://www.sse.com.cn/assortment/stock/list/share/"
)

type SHResponse struct {
    Data  []SHResponseData `json:"result"`
    Error string           `json:"csrcCode"`
}

type SHResponseData struct {
    CompanySimpleName string `json:"COMPANY_ABBR"`
    StockCode         string `json:"COMPANY_CODE"`
    ListingDate       string `json:"LISTING_DATE"`
}

func (s SHCompanyCollector) String() string {
    return "SHCompanyCollector"
}

func (s SHCompanyCollector) GetStockExchange() int {
    return enums.SH
}

func (s SHCompanyCollector) FetchAll(conf config.StockConfig) []model.Company {
    logger.Infow("收集上交所公司信息,开始.", "stockExchangeCode", s.GetStockExchange(), "configInfo", conf)
    result := make([]model.Company, 0)
    plates := enums.GetByStockExchange(conf)
    for _, plate := range plates {
        result = append(result, SHGetPlateData(conf, plate)...)
    }
    logger.Infow("收集上交所公司信息,结束.", "stockExchangeCode", s.GetStockExchange(), "configInfo", conf, "length", len(result))
    return result
}

//获取每个板块的数据
func SHGetPlateData(conf config.StockConfig, plate enums.PlateEnum) []model.Company {
    logger.Infow("收集上交所公司信息,收集指定板块信息,开始.", "stockExchangeCode", conf.StockExchangeCode, "plate", plate)
    result := make([]model.Company, 0)
    page := 1
    for pageData := SHReadPageData(conf, page, plate); pageData != nil; {
        result = append(result, pageData...)
        page = page + 1
        pageData = SHReadPageData(conf, page, plate)
    }
    logger.Infow("收集上交所公司信息,收集指定板块信息,结束.", "stockExchangeCode", conf.StockExchangeCode, "plate", plate)
    return result
}

//读取每页的数据
func SHReadPageData(conf config.StockConfig, page int, plate enums.PlateEnum) []model.Company {
    time.Sleep(time.Millisecond * 500)
    client := &http.Client{}
    pageStr := strconv.Itoa(page)
    requestUrl := conf.CompanyInfoUrl + "&pageHelp.beginPage=" + pageStr + "&pageHelp.pageNo=" + pageStr + "&stockType=" + plate.Tab + "&_=" + strconv.Itoa(rand.Int())
    logger.Infow("获取上交所公司列表.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl)
    request, err := http.NewRequest("GET", requestUrl, nil)
    if err != nil {
        logger.Errorw("获取上交所公司列表构造请求参数异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
        return nil
    }
    request.Header.Add("Cookie", COOKIE)
    request.Header.Add("User-Agent", USERAGENT)
    request.Header.Add("Referer", Referer)
    response, err := client.Do(request)
    if nil != err {
        logger.Errorw("获取上交所公司列表数据异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
        return nil
    }
    responseDataByte, err := ioutil.ReadAll(response.Body)
    if nil != err {
        logger.Errorw("读取上交所公司列表数据异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
        return nil
    }
    responseDataPoint := &SHResponse{}
    err = json.Unmarshal(responseDataByte, &responseDataPoint)
    if err != nil {
        logger.Errorw("读取上交所公司列表,解析数据出现异常.", "stockExchange", plate.StockExchange, "plate", plate.Tab, "url", requestUrl, "err", err)
        return nil
    }
    if nil == responseDataPoint || len(responseDataPoint.Data) == 0 {
        return nil
    }
    return SHResponseToCompanyMapper(*responseDataPoint, plate)
}

func SHResponseToCompanyMapper(response SHResponse, plate enums.PlateEnum) []model.Company {
    result := make([]model.Company, 0)
    for _, d := range response.Data {
        company := model.Company{
            StockExchange: plate.StockExchange,
            Code:          d.StockCode,
            Plate:         plate.Code,
            ShortName:     d.CompanySimpleName,
            FullName:      "-",
            IndustryCode:  "-",
            IndustryName:  "-",
            ListingDate:   d.ListingDate,
        }
        result = append(result, company)
    }
    logger.Infow("数据转换", "length", len(response.Data), "resultLength", len(result))
    return result
}
