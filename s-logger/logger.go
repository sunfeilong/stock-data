package s_logger

import (
    "../tool"
    "go.uber.org/zap"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

var zapC zap.Config

const maxLevel int = 6

//配置文件名字
const (
    logConfigFileName string = "log.yml"
)

//加载配置文件
func init() {
    zapConfig, err := ReadZapConfig()
    if nil != err {
        panic(err)
    }
    zapC = *zapConfig
}

func New() *zap.SugaredLogger {
    logger, err := zapC.Build()
    if nil != err {
        panic(err)
    }
    return logger.Sugar()
}

func ReadZapConfig() (*zap.Config, error) {
    config := zap.Config{}
    filePath, eil := tool.GetConfigFilePath(logConfigFileName, maxLevel, maxLevel)
    log.Println("配置文件路径: ", filePath)
    log.Println("读取Zap配置文件.", "filePath", filePath)
    if nil != eil {
        log.Println("读取Zap配置文件,未找到配置文件.", "filePath: ", filePath, "error: ", eil)
        return &config, eil
    }
    yamlFile, eil := ioutil.ReadFile(filePath)
    if nil != eil {
        log.Println("读取Zap配置文件,读取配置信息出错.", "filePath: ", filePath, "error: ", eil)
        return &config, eil
    }
    eil = yaml.Unmarshal(yamlFile, &config)
    if nil != eil {
        log.Println("读取Zap配置文件,解析配置文件出错.", "filePath: ", filePath, "error: ", eil)
        return &config, eil
    }
    return &config, nil
}
