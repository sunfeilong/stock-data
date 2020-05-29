package s_logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
    "time"
)

func NewEncoderConfig() zapcore.EncoderConfig {
    return zapcore.EncoderConfig{
        TimeKey:        "Time",
        LevelKey:       "Level",
        NameKey:        "Name",
        CallerKey:      "Caller",
        MessageKey:     "Message",
        StacktraceKey:  "Stack",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.CapitalLevelEncoder,
        EncodeTime:     TimeEncoder,
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }
}

func New() *zap.SugaredLogger {

    GTEError := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl >= zapcore.ErrorLevel
    })

    GETDebug := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl > zapcore.DebugLevel
    })

    infoOutput := zapcore.AddSync(&lumberjack.Logger{
        Filename:   "log-info.log",
        MaxSize:    1,
        MaxBackups: 3,
        MaxAge:     28,
    })

    errorOutput := zapcore.AddSync(&lumberjack.Logger{
        Filename:   "log-error.log",
        MaxSize:    1,
        MaxBackups: 3,
        MaxAge:     28,
    })

    console := zapcore.Lock(os.Stdout)
    prodEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
    devEncoder := zapcore.NewConsoleEncoder(NewEncoderConfig())

    core := zapcore.NewTee(
        zapcore.NewCore(devEncoder, console, GETDebug),

        zapcore.NewCore(prodEncoder, infoOutput, GETDebug),
        zapcore.NewCore(prodEncoder, errorOutput, GTEError),
    )
    logger := zap.New(core, zap.AddCaller())
    return logger.Sugar()
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
