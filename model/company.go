package model

//股票公司
type Company struct {
    StockExchange int    `json:"stock_exchange"`
    Code          string `json:"code"`
    Plate         string `json:"plate"`
    ShortName     string `json:"short_name"`
    FullName      string `json:"full_name"`
    IndustryCode  string `json:"industry_code"`
    IndustryName  string `json:"industry_name"`
    ListingDate   string `json:"listing_date"`
}
