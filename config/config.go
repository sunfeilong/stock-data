package config

import (
    "../enums"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

//配置文件名字
const fileName string = "../configs.yml"

type Cache interface {
    //获取配置信息
    getConfig(exchange enums.StockExchange) (Config, error)
}

type YamlConfig struct {
}

//配置信息定义
type Config struct {
    StockExchangeCode enums.StockExchange
    StockExchange     string
    CompanyInfoUrl    string
    StockInfoUrl      string
    RealTimeInfoUrl   string
}

//配置信息列表
type Configs struct {
    Configs []Config
}

//配置信息缓存
var configMap map[enums.StockExchange]Config

//获取指定股票交易所的配置信息
func (yamlConfig YamlConfig) getConfig(sc enums.StockExchange) (Config, error) {
    if configMap != nil {
        log.Println("配置文件信息已经加载,配置信息: ", configMap)
        return configMap[sc], nil
    }
    log.Println("配置文件信息不存在,加载配置信息开始")
    err := yamlConfig.loadConfig()
    log.Println("配置文件信息不存在,加载配置信息结束. 配置信息: ", configMap)
    if nil != err {
        log.Fatal("加载配置文件出错,fileName: ", fileName, ", error info: ", err.Error())
        return *new(Config), err
    }
    return configMap[sc], nil
}

//加载配置文件
func (yamlConfig YamlConfig) loadConfig() error {
    configMap = make(map[enums.StockExchange]Config)
    configs := Configs{}
    yamlFIle, eil := ioutil.ReadFile(fileName)
    if nil != eil {
        log.Fatal("读取配置文件出错,fileName: ", fileName, ", error info: ", eil.Error())
        return eil
    }
    eil = yaml.Unmarshal(yamlFIle, &configs)
    if nil != eil {
        log.Fatal("解析配置文件出错,fileName: ", fileName, ", error info: ", eil.Error())
        return eil
    }
    cs := configs.Configs

    for i := range cs {
        configMap[cs[i].StockExchangeCode] = cs[i]
    }
    return nil
}
