package s_logger

import (
    "../tool"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
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
    //logger, err := zapC.Build()
    highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl >= zapcore.ErrorLevel
    })

    lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl <= zapcore.ErrorLevel
    })

    infoFile, _ := os.OpenFile(zapC.OutputPaths[1], os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    errorFile, _ := os.OpenFile(zapC.ErrorOutputPaths[1], os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    infoOutput := zapcore.AddSync(infoFile)
    errorOutput := zapcore.AddSync(errorFile)

    console := zapcore.Lock(os.Stdout)

    prodEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
    devEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

    core := zapcore.NewTee(
        zapcore.NewCore(devEncoder, console, lowPriority),

        zapcore.NewCore(prodEncoder, infoOutput, lowPriority),
        zapcore.NewCore(prodEncoder, errorOutput, highPriority),
    )
    logger := zap.New(core)
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
