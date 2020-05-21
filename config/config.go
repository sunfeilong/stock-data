package config

import (
    "../enums"
    "../errors"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
    "sync"
)

var mu sync.Mutex

const maxLevel int = 6

//配置文件名字
const fileName string = "configs.yml"

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
func (yamlConfig YamlConfig) GetConfig(sc enums.StockExchange) (Config, error) {
    if configMap != nil {
        log.Println("配置文件信息已经加载,配置信息: ", configMap)
        return configMap[sc], nil
    }
    mu.Lock()
    if configMap != nil {
        log.Println("配置文件信息已经加载,配置信息: ", configMap)
        return configMap[sc], nil
    }
    log.Println("配置文件信息没有初始化,加载配置信息开始")
    err := yamlConfig.loadConfig()
    log.Println("配置文件信息没有初始化,加载配置信息结束. 配置信息: ", configMap)
    mu.Unlock()
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
    filePath, eil := getConfigFilePath(fileName, maxLevel)
    log.Println("配置文件路径: ", filePath)
    if nil != eil {
        log.Fatal("读取配置文件出错,未找到配置文件.filePath: ", filePath, ", error info: ", eil.Error())
        return eil
    }
    yamlFile, eil := ioutil.ReadFile(filePath)
    if nil != eil {
        log.Fatal("读取配置文件出错,filePath: ", filePath, ", error info: ", eil.Error())
        return eil
    }
    eil = yaml.Unmarshal(yamlFile, &configs)
    if nil != eil {
        log.Fatal("解析配置文件出错,filePath: ", filePath, ", error info: ", eil.Error())
        return eil
    }
    for _, config := range configs.Configs {
        configMap[config.StockExchangeCode] = config
    }
    return nil
}

func getConfigFilePath(fileName string, level int) (string, error) {
    log.Println("寻找配置文件，配置文件路径: ", fileName, " 向上遍历层级: ", maxLevel-level)
    _, err := os.Open(fileName)
    if level < 0 {
        log.Fatal("寻找配置文件，配置文件路径: ", fileName)
        return "", errors.StockDataError{"未找到配置文件"}
    }
    if nil == err {
        return fileName, nil
    } else {
        return getConfigFilePath("../"+fileName, level-1)
    }
}
