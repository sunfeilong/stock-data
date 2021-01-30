## 项目说明

收集上市公司股票价格数据，项目只支持收集国内交易所的当天的股票价格数据。

用于闭市之后收集当天的股票价格数据。

**配置 project_config.yml:**

* `datasavepath`：数据文件保存路径，需要配置。

**数据文件命名:**

* 公司数据: `company-年-月-日.json` 

    ```json
    [
      {
        "stock_exchange": 100,
        "code": "000001",
        "plate": "100",
        "short_name": "平安银行",
        "full_name": "平安银行股份有限公司",
        "industry_code": "J",
        "industry_name": "金融业"
      },
      ...
     ]
    ```

* 股票数据: `data-年-月-日.json`

    ```json
    [
      {
        "stock_exchange": 100,
        "code": "000001",
        "plate": "100",
        "data": [
          {
            "date": "2020-06-11 09:30",
            "price": "13.38"
          },
          ...
        ]
      },
      {
        "stock_exchange": 100,
        "code": "000002",
        "plate": "100",
        "data": [
          {
            "date": "2020-06-11 09:30",
            "price": "13.38"
          },
          ...
        ]
      }
      ...
    ]
    ```

### 运行项目

1. 直接运行

    ```
    git clone https://github.com/sunfeilong/stock-data.git
    cd stock-data
    go run stock_data.go
    ```
    
1. 打包之后运行 

    ```
    git clone https://github.com/sunfeilong/stock-data.git
    cd stock-data
    go build
    ./stock
    ```

### 注意事项

为了不影响数据源网站正常运行，使用时请不要删除限速的代码。

## 功能列表

### 已完成

* [深圳证券交易所](http://www.szse.cn/)数据收集。
* [上海证券交易所](http://www.sse.com.cn/)数据收集。
* [香港交易所](https://sc.hkex.com.hk/TuniS/www.hkex.com.hk/?sc_lang=zh-cn)数据收集。

### 正在做

* 无

### 下一步

* 无

### 已废弃的功能

* 每天16:30定时执行数据收集任务。收集完成之后自动推送数据到该仓库。

    废弃原因：随着收集次数的变多，收集的数据文件慢慢变大，代码仓库开始变的很大，拉取代码开始变得很慢，而且总有一天本地磁盘会不够用，所以废弃了该功能。

