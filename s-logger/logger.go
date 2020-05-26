package s_logger

import (
    "../tool"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
    "time"
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

func NewEncoderConfig() zapcore.EncoderConfig {
    return zapcore.EncoderConfig{
        // Keys can be anything except the empty string.
        TimeKey:        "T",
        LevelKey:       "L",
        NameKey:        "N",
        CallerKey:      "C",
        MessageKey:     "M",
        StacktraceKey:  "S",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalLevelEncoder,
        EncodeTime:     TimeEncoder,
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }
}

func New() *zap.SugaredLogger {
    //logger, err := zapC.Build()

    GTEError := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl >= zapcore.ErrorLevel
    })

    GETDebug := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl > zapcore.DebugLevel
    })

    infoOutput := zapcore.AddSync(&lumberjack.Logger{
        Filename:   zapC.OutputPaths[1],
        MaxSize:    1, // megabytes
        MaxBackups: 3,
        MaxAge:     28, // days
    })

    errorOutput := zapcore.AddSync(&lumberjack.Logger{
        Filename:   zapC.ErrorOutputPaths[1],
        MaxSize:    1, // megabytes
        MaxBackups: 3,
        MaxAge:     28, // days
    })

    console := zapcore.Lock(os.Stdout)

    prodEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
    devEncoder := zapcore.NewConsoleEncoder(NewEncoderConfig())

    core := zapcore.NewTee(
        zapcore.NewCore(devEncoder, console, GETDebug),

        zapcore.NewCore(prodEncoder, infoOutput, GETDebug),
        zapcore.NewCore(prodEncoder, errorOutput, GTEError),
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

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
