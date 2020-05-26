package config

import (
    "../enums"
    "../s-logger"
    "../tool"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "sync"
)

const maxLevel int = 6

//配置文件名字
const (
    defaultConfigFileName string = "configs.yml"
    logConfigFileName     string = "log.yml"
)

type YamlConfig struct {
    sync.Mutex
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

var logger = s_logger.New()

//获取指定股票交易所的配置信息
func (yamlConfig YamlConfig) GetConfig(sc enums.StockExchange) (Config, error) {
    if configMap != nil {
        logger.Infow("配置文件信息已经加载.", "配置信息:", configMap)
        return configMap[sc], nil
    }
    yamlConfig.Lock()
    if configMap != nil {
        logger.Infow("配置文件信息已经加载.", "配置信息:", configMap)
        return configMap[sc], nil
    }
    logger.Infow("配置文件信息没有初始化,加载配置信息开始.")
    err := loadConfig()
    logger.Infow("配置文件信息没有初始化,加载配置信息结束.", "配置信息", configMap)
    yamlConfig.Unlock()
    if nil != err {
        logger.Errorw("加载配置文件出错.", "defaultConfigFileName", defaultConfigFileName, "error", err)
        return *new(Config), err
    }
    return configMap[sc], nil
}

//加载配置文件
func loadConfig() error {
    configMap = make(map[enums.StockExchange]Config)
    configs := Configs{}
    filePath, eil := tool.GetConfigFilePath(defaultConfigFileName, maxLevel, maxLevel)
    logger.Infow("加载配置文件.", "配置文件路径", filePath)
    if nil != eil {
        logger.Errorw("读取配置文件出错,未找到配置文件..", "filePath", filePath, "error", eil)
        return eil
    }
    yamlFile, eil := ioutil.ReadFile(filePath)
    if nil != eil {
        logger.Errorw("读取配置信息出错.", "filePath", filePath, "error", eil)
        return eil
    }
    eil = yaml.Unmarshal(yamlFile, &configs)
    if nil != eil {
        logger.Errorw("解析配置文件出错.", "filePath", filePath, "error", eil)
        return eil
    }
    for _, config := range configs.Configs {
        configMap[config.StockExchangeCode] = config
    }
    return nil
}
